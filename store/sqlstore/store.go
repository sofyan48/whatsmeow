// Copyright (c) 2022 Tulir Asokan
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package sqlstore contains an SQL-backed implementation of the interfaces in the store package.
package sqlstore

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/kiriminaja/kaj-notification-srvc/src/pkg/whatsapp/whatsmeow/store"
	"github.com/kiriminaja/kaj-notification-srvc/src/pkg/whatsapp/whatsmeow/types"
	"github.com/kiriminaja/kaj-notification-srvc/src/pkg/whatsapp/whatsmeow/util/keys"
)

// ErrInvalidLength is returned by some database getters if the database returned a byte array with an unexpected length.
// This should be impossible, as the database schema contains CHECK()s for all the relevant columns.
var ErrInvalidLength = errors.New("database returned byte array with illegal length")

// PostgresArrayWrapper is a function to wrap array values before passing them to the sql package.
//
// When using github.com/lib/pq, you should set
//
//	whatsmeow.PostgresArrayWrapper = pq.Array
var PostgresArrayWrapper func(interface{}) interface {
	driver.Valuer
	sql.Scanner
}

type SQLStore struct {
	*Container
	JID string

	preKeyLock sync.Mutex

	contactCache     map[types.JID]*types.ContactInfo
	contactCacheLock sync.Mutex
}

// NewSQLStore creates a new SQLStore with the given database container and user JID.
// It contains implementations of all the different stores in the store package.
//
// In general, you should use Container.NewDevice or Container.GetDevice instead of this.
func NewSQLStore(c *Container, jid types.JID) *SQLStore {
	return &SQLStore{
		Container:    c,
		JID:          jid.String(),
		contactCache: make(map[types.JID]*types.ContactInfo),
	}
}

var _ store.IdentityStore = (*SQLStore)(nil)
var _ store.SessionStore = (*SQLStore)(nil)
var _ store.PreKeyStore = (*SQLStore)(nil)
var _ store.SenderKeyStore = (*SQLStore)(nil)
var _ store.AppStateSyncKeyStore = (*SQLStore)(nil)
var _ store.AppStateStore = (*SQLStore)(nil)
var _ store.ContactStore = (*SQLStore)(nil)

const (
	putIdentityQuery = `INSERT INTO whatsmeow_identity_keys (our_jid, their_id, identity)
	VALUES (?, ?, ?)
	ON DUPLICATE KEY UPDATE
		identity = VALUES(identity)`
	deleteAllIdentitiesQuery = `DELETE FROM whatsmeow_identity_keys WHERE our_jid=? AND their_id LIKE ?`
	deleteIdentityQuery      = `DELETE FROM whatsmeow_identity_keys WHERE our_jid=? AND their_id=?`
	getIdentityQuery         = `SELECT identity FROM whatsmeow_identity_keys WHERE our_jid=? AND their_id=?`
)

func (s *SQLStore) PutIdentity(address string, key [32]byte) error {
	_, err := s.db.Exec(putIdentityQuery, s.JID, address, key[:])
	return err
}

func (s *SQLStore) DeleteAllIdentities(phone string) error {
	_, err := s.db.Exec(deleteAllIdentitiesQuery, s.JID, phone+":%")
	return err
}

func (s *SQLStore) DeleteIdentity(address string) error {
	_, err := s.db.Exec(deleteAllIdentitiesQuery, s.JID, address)
	return err
}

func (s *SQLStore) IsTrustedIdentity(address string, key [32]byte) (bool, error) {
	var existingIdentity []byte
	err := s.db.QueryRow(getIdentityQuery, s.JID, address).Scan(&existingIdentity)
	if errors.Is(err, sql.ErrNoRows) {
		// Trust if not known, it'll be saved automatically later
		return true, nil
	} else if err != nil {
		return false, err
	} else if len(existingIdentity) != 32 {
		return false, ErrInvalidLength
	}
	return *(*[32]byte)(existingIdentity) == key, nil
}

const (
	getSessionQuery = `SELECT session FROM whatsmeow_sessions WHERE our_jid=? AND their_id=?`
	hasSessionQuery = `SELECT true FROM whatsmeow_sessions WHERE our_jid=? AND their_id=?`
	putSessionQuery = `INSERT INTO whatsmeow_sessions (our_jid, their_id, session)
	VALUES (?, ?, ?)
	ON DUPLICATE KEY UPDATE
		session = VALUES(session)`
	deleteAllSessionsQuery = `DELETE FROM whatsmeow_sessions WHERE our_jid=? AND their_id LIKE ?`
	deleteSessionQuery     = `DELETE FROM whatsmeow_sessions WHERE our_jid=? AND their_id=?`
)

func (s *SQLStore) GetSession(address string) (session []byte, err error) {
	err = s.db.QueryRow(getSessionQuery, s.JID, address).Scan(&session)
	if errors.Is(err, sql.ErrNoRows) {
		err = nil
	}
	return
}

func (s *SQLStore) HasSession(address string) (has bool, err error) {
	err = s.db.QueryRow(hasSessionQuery, s.JID, address).Scan(&has)
	if errors.Is(err, sql.ErrNoRows) {
		err = nil
	}
	return
}

func (s *SQLStore) PutSession(address string, session []byte) error {
	_, err := s.db.Exec(putSessionQuery, s.JID, address, session)
	return err
}

func (s *SQLStore) DeleteAllSessions(phone string) error {
	_, err := s.db.Exec(deleteAllSessionsQuery, s.JID, phone+":%")
	return err
}

func (s *SQLStore) DeleteSession(address string) error {
	_, err := s.db.Exec(deleteSessionQuery, s.JID, address)
	return err
}

const (
	getLastPreKeyIDQuery        = `SELECT MAX(key_id) FROM whatsmeow_pre_keys WHERE jid=?`
	insertPreKeyQuery           = `INSERT INTO whatsmeow_pre_keys (jid, key_id, ` + "`key`" + `, uploaded) VALUES (?, ?, ?, ?)`
	getUnuploadedPreKeysQuery   = `SELECT key_id, ` + "`key`" + ` FROM whatsmeow_pre_keys WHERE jid=? AND uploaded=false ORDER BY key_id LIMIT ?`
	getPreKeyQuery              = `SELECT key_id, ` + "`key`" + ` FROM whatsmeow_pre_keys WHERE jid=? AND key_id=?`
	deletePreKeyQuery           = `DELETE FROM whatsmeow_pre_keys WHERE jid=? AND key_id=?`
	markPreKeysAsUploadedQuery  = `UPDATE whatsmeow_pre_keys SET uploaded=true WHERE jid=? AND key_id<=?`
	getUploadedPreKeyCountQuery = `SELECT COUNT(*) FROM whatsmeow_pre_keys WHERE jid=? AND uploaded=true`
)

func (s *SQLStore) genOnePreKey(id uint32, markUploaded bool) (*keys.PreKey, error) {
	key := keys.NewPreKey(id)
	_, err := s.db.Exec(insertPreKeyQuery, s.JID, key.KeyID, key.Priv[:], markUploaded)
	return key, err
}

func (s *SQLStore) getNextPreKeyID() (uint32, error) {
	var lastKeyID sql.NullInt32
	err := s.db.QueryRow(getLastPreKeyIDQuery, s.JID).Scan(&lastKeyID)
	if err != nil {
		return 0, fmt.Errorf("failed to query next prekey ID: %w", err)
	}
	return uint32(lastKeyID.Int32) + 1, nil
}

func (s *SQLStore) GenOnePreKey() (*keys.PreKey, error) {
	s.preKeyLock.Lock()
	defer s.preKeyLock.Unlock()
	nextKeyID, err := s.getNextPreKeyID()
	if err != nil {
		return nil, err
	}
	return s.genOnePreKey(nextKeyID, true)
}

func (s *SQLStore) GetOrGenPreKeys(count uint32) ([]*keys.PreKey, error) {
	s.preKeyLock.Lock()
	defer s.preKeyLock.Unlock()

	res, err := s.db.Query(getUnuploadedPreKeysQuery, s.JID, count)
	if err != nil {
		return nil, fmt.Errorf("failed to query existing prekeys: %w", err)
	}
	newKeys := make([]*keys.PreKey, count)
	var existingCount uint32
	for res.Next() {
		var key *keys.PreKey
		key, err = scanPreKey(res)
		if err != nil {
			return nil, err
		} else if key != nil {
			newKeys[existingCount] = key
			existingCount++
		}
	}

	if existingCount < uint32(len(newKeys)) {
		var nextKeyID uint32
		nextKeyID, err = s.getNextPreKeyID()
		if err != nil {
			return nil, err
		}
		for i := existingCount; i < count; i++ {
			newKeys[i], err = s.genOnePreKey(nextKeyID, false)
			if err != nil {
				return nil, fmt.Errorf("failed to generate prekey: %w", err)
			}
			nextKeyID++
		}
	}

	return newKeys, nil
}

func scanPreKey(row scannable) (*keys.PreKey, error) {
	var priv []byte
	var id uint32
	err := row.Scan(&id, &priv)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	} else if len(priv) != 32 {
		return nil, ErrInvalidLength
	}
	return &keys.PreKey{
		KeyPair: *keys.NewKeyPairFromPrivateKey(*(*[32]byte)(priv)),
		KeyID:   id,
	}, nil
}

func (s *SQLStore) GetPreKey(id uint32) (*keys.PreKey, error) {
	return scanPreKey(s.db.QueryRow(getPreKeyQuery, s.JID, id))
}

func (s *SQLStore) RemovePreKey(id uint32) error {
	_, err := s.db.Exec(deletePreKeyQuery, s.JID, id)
	return err
}

func (s *SQLStore) MarkPreKeysAsUploaded(upToID uint32) error {
	_, err := s.db.Exec(markPreKeysAsUploadedQuery, s.JID, upToID)
	return err
}

func (s *SQLStore) UploadedPreKeyCount() (count int, err error) {
	err = s.db.QueryRow(getUploadedPreKeyCountQuery, s.JID).Scan(&count)
	return
}

const (
	getSenderKeyQuery = `SELECT sender_key FROM whatsmeow_sender_keys WHERE our_jid=? AND chat_id=? AND sender_id=?`
	putSenderKeyQuery = `
	INSERT INTO whatsmeow_sender_keys (our_jid, chat_id, sender_id, sender_key)
	VALUES (?, ?, ?, ?)
	ON DUPLICATE KEY UPDATE
		sender_key = VALUES(sender_key)
	`
)

func (s *SQLStore) PutSenderKey(group, user string, session []byte) error {
	_, err := s.db.Exec(putSenderKeyQuery, s.JID, group, user, session)
	return err
}

func (s *SQLStore) GetSenderKey(group, user string) (key []byte, err error) {
	err = s.db.QueryRow(getSenderKeyQuery, s.JID, group, user).Scan(&key)
	if errors.Is(err, sql.ErrNoRows) {
		err = nil
	}
	return
}

const (
	putAppStateSyncKeyQuery = `INSERT INTO whatsmeow_app_state_sync_keys (jid, key_id, key_data, timestamp, fingerprint)
	VALUES (?, ?, ?, ?, ?)
	ON DUPLICATE KEY UPDATE
		key_data = CASE
			WHEN VALUES(timestamp) > timestamp THEN VALUES(key_data)
			ELSE key_data
		END,
		timestamp = CASE
			WHEN VALUES(timestamp) > timestamp THEN VALUES(timestamp)
			ELSE timestamp
		END,
		fingerprint = CASE
			WHEN VALUES(timestamp) > timestamp THEN VALUES(fingerprint)
			ELSE fingerprint
		END`
	getAppStateSyncKeyQuery         = `SELECT key_data, timestamp, fingerprint FROM whatsmeow_app_state_sync_keys WHERE jid=? AND key_id=?`
	getLatestAppStateSyncKeyIDQuery = `SELECT key_id FROM whatsmeow_app_state_sync_keys WHERE jid=? ORDER BY timestamp DESC LIMIT 1`
)

func (s *SQLStore) PutAppStateSyncKey(id []byte, key store.AppStateSyncKey) error {
	_, err := s.db.Exec(putAppStateSyncKeyQuery, s.JID, id, key.Data, key.Timestamp, key.Fingerprint)
	return err
}

func (s *SQLStore) GetAppStateSyncKey(id []byte) (*store.AppStateSyncKey, error) {
	var key store.AppStateSyncKey
	err := s.db.QueryRow(getAppStateSyncKeyQuery, s.JID, id).Scan(&key.Data, &key.Timestamp, &key.Fingerprint)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return &key, err
}

func (s *SQLStore) GetLatestAppStateSyncKeyID() ([]byte, error) {
	var keyID []byte
	err := s.db.QueryRow(getLatestAppStateSyncKeyIDQuery, s.JID).Scan(&keyID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return keyID, err
}

const (
	putAppStateVersionQuery = `INSERT INTO whatsmeow_app_state_version (jid, name, version, hash)
	VALUES (?, ?, ?, ?)
	ON DUPLICATE KEY UPDATE
		version = VALUES(version),
		hash = VALUES(hash)`
	getAppStateVersionQuery                 = `SELECT version, hash FROM whatsmeow_app_state_version WHERE jid=? AND name=?`
	deleteAppStateVersionQuery              = `DELETE FROM whatsmeow_app_state_version WHERE jid=? AND name=?`
	putAppStateMutationMACsQuery            = `INSERT INTO whatsmeow_app_state_mutation_macs (jid, name, version, index_mac, value_mac) VALUES `
	deleteAppStateMutationMACsQueryPostgres = `DELETE FROM whatsmeow_app_state_mutation_macs WHERE jid=? AND name=? AND index_mac=ANY(?::bytea[])`
	deleteAppStateMutationMACsQueryGeneric  = `DELETE FROM whatsmeow_app_state_mutation_macs WHERE jid=? AND name=? AND index_mac IN `
	getAppStateMutationMACQuery             = `SELECT value_mac FROM whatsmeow_app_state_mutation_macs WHERE jid=? AND name=? AND index_mac=? ORDER BY version DESC LIMIT 1`
)

func (s *SQLStore) PutAppStateVersion(name string, version uint64, hash [128]byte) error {
	_, err := s.db.Exec(putAppStateVersionQuery, s.JID, name, version, hash[:])
	return err
}

func (s *SQLStore) GetAppStateVersion(name string) (version uint64, hash [128]byte, err error) {
	var uncheckedHash []byte
	err = s.db.QueryRow(getAppStateVersionQuery, s.JID, name).Scan(&version, &uncheckedHash)
	if errors.Is(err, sql.ErrNoRows) {
		// version will be 0 and hash will be an empty array, which is the correct initial state
		err = nil
	} else if err != nil {
		// There's an error, just return it
	} else if len(uncheckedHash) != 128 {
		// This shouldn't happen
		err = ErrInvalidLength
	} else {
		// No errors, convert hash slice to array
		hash = *(*[128]byte)(uncheckedHash)
	}
	return
}

func (s *SQLStore) DeleteAppStateVersion(name string) error {
	_, err := s.db.Exec(deleteAppStateVersionQuery, s.JID, name)
	return err
}

type execable interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

func (s *SQLStore) putAppStateMutationMACs(tx execable, name string, version uint64, mutations []store.AppStateMutationMAC) error {
	values := make([]interface{}, 3+len(mutations)*2)
	queryParts := make([]string, len(mutations))
	values[0] = s.JID
	values[1] = name
	values[2] = version
	placeholderSyntax := "(?, ?, ?, ?, ?)"
	if s.dialect == "sqlite3" {
		placeholderSyntax = "(?1, ?2, ?3, ?%d, ?%d)"
	}
	for i, mutation := range mutations {
		baseIndex := 3 + i*2
		values[baseIndex] = mutation.IndexMAC
		values[baseIndex+1] = mutation.ValueMAC
		queryParts[i] = fmt.Sprintf(placeholderSyntax, baseIndex+1, baseIndex+2)
	}
	_, err := tx.Exec(putAppStateMutationMACsQuery+strings.Join(queryParts, ","), values...)
	return err
}

const mutationBatchSize = 400

func (s *SQLStore) PutAppStateMutationMACs(name string, version uint64, mutations []store.AppStateMutationMAC) error {
	if len(mutations) > mutationBatchSize {
		tx, err := s.db.Begin()
		if err != nil {
			return fmt.Errorf("failed to start transaction: %w", err)
		}
		for i := 0; i < len(mutations); i += mutationBatchSize {
			var mutationSlice []store.AppStateMutationMAC
			if len(mutations) > i+mutationBatchSize {
				mutationSlice = mutations[i : i+mutationBatchSize]
			} else {
				mutationSlice = mutations[i:]
			}
			err = s.putAppStateMutationMACs(tx, name, version, mutationSlice)
			if err != nil {
				_ = tx.Rollback()
				return err
			}
		}
		err = tx.Commit()
		if err != nil {
			return fmt.Errorf("failed to commit transaction: %w", err)
		}
		return nil
	} else if len(mutations) > 0 {
		return s.putAppStateMutationMACs(s.db, name, version, mutations)
	}
	return nil
}

func (s *SQLStore) DeleteAppStateMutationMACs(name string, indexMACs [][]byte) (err error) {
	if len(indexMACs) == 0 {
		return
	}
	if s.dialect == "postgres" && PostgresArrayWrapper != nil {
		_, err = s.db.Exec(deleteAppStateMutationMACsQueryPostgres, s.JID, name, PostgresArrayWrapper(indexMACs))
	} else {
		args := make([]interface{}, 2+len(indexMACs))
		args[0] = s.JID
		args[1] = name
		queryParts := make([]string, len(indexMACs))
		for i, item := range indexMACs {
			args[2+i] = item
			queryParts[i] = fmt.Sprintf("$%d", i+3)
		}
		_, err = s.db.Exec(deleteAppStateMutationMACsQueryGeneric+"("+strings.Join(queryParts, ",")+")", args...)
	}
	return
}

func (s *SQLStore) GetAppStateMutationMAC(name string, indexMAC []byte) (valueMAC []byte, err error) {
	err = s.db.QueryRow(getAppStateMutationMACQuery, s.JID, name, indexMAC).Scan(&valueMAC)
	if errors.Is(err, sql.ErrNoRows) {
		err = nil
	}
	return
}

const (
	putContactNameQuery = `INSERT INTO whatsmeow_contacts (our_jid, their_jid, first_name, full_name)
	VALUES (?, ?, ?, ?)
	ON DUPLICATE KEY UPDATE
		first_name = VALUES(first_name),
		full_name = VALUES(full_name)`

	putManyContactNamesQuery = `
		INSERT INTO whatsmeow_contacts (our_jid, their_jid, first_name, full_name)
		VALUES %s
		ON DUPLICATE KEY UPDATE
			first_name = VALUES(first_name),
			full_name = VALUES(full_name)`
	putPushNameQuery = `INSERT INTO whatsmeow_contacts (our_jid, their_jid, push_name)
	VALUES (?, ?, ?)
	ON DUPLICATE KEY UPDATE
		push_name = VALUES(push_name)`

	putBusinessNameQuery = `
	INSERT INTO whatsmeow_contacts (our_jid, their_jid, business_name)
	VALUES (?, ?, ?)
	ON DUPLICATE KEY UPDATE
		business_name = VALUES(business_name)`

	getContactQuery = `
		SELECT first_name, full_name, push_name, business_name FROM whatsmeow_contacts WHERE our_jid=? AND their_jid=?
	`
	getAllContactsQuery = `
		SELECT their_jid, first_name, full_name, push_name, business_name FROM whatsmeow_contacts WHERE our_jid=?
	`
)

func (s *SQLStore) PutPushName(user types.JID, pushName string) (bool, string, error) {
	s.contactCacheLock.Lock()
	defer s.contactCacheLock.Unlock()

	cached, err := s.getContact(user)
	if err != nil {
		return false, "", err
	}
	if cached.PushName != pushName {
		_, err = s.db.Exec(putPushNameQuery, s.JID, user, pushName)
		if err != nil {
			return false, "", err
		}
		previousName := cached.PushName
		cached.PushName = pushName
		cached.Found = true
		return true, previousName, nil
	}
	return false, "", nil
}

func (s *SQLStore) PutBusinessName(user types.JID, businessName string) (bool, string, error) {
	s.contactCacheLock.Lock()
	defer s.contactCacheLock.Unlock()

	cached, err := s.getContact(user)
	if err != nil {
		return false, "", err
	}
	if cached.BusinessName != businessName {
		_, err = s.db.Exec(putBusinessNameQuery, s.JID, user, businessName)
		if err != nil {
			return false, "", err
		}
		previousName := cached.BusinessName
		cached.BusinessName = businessName
		cached.Found = true
		return true, previousName, nil
	}
	return false, "", nil
}

func (s *SQLStore) PutContactName(user types.JID, firstName, fullName string) error {
	s.contactCacheLock.Lock()
	defer s.contactCacheLock.Unlock()

	cached, err := s.getContact(user)
	if err != nil {
		return err
	}
	if cached.FirstName != firstName || cached.FullName != fullName {
		_, err = s.db.Exec(putContactNameQuery, s.JID, user, firstName, fullName)
		if err != nil {
			return err
		}
		cached.FirstName = firstName
		cached.FullName = fullName
		cached.Found = true
	}
	return nil
}

const contactBatchSize = 300

func (s *SQLStore) putContactNamesBatch(tx execable, contacts []store.ContactEntry) error {
	values := make([]interface{}, 1, 1+len(contacts)*3)
	queryParts := make([]string, 0, len(contacts))
	values[0] = s.JID
	placeholderSyntax := "(?, ?, ?, ?)"
	if s.dialect == "sqlite3" {
		placeholderSyntax = "(?1, ?%d, ?%d, ?%d)"
	}
	i := 0
	handledContacts := make(map[types.JID]struct{}, len(contacts))
	for _, contact := range contacts {
		if contact.JID.IsEmpty() {
			s.log.Warnf("Empty contact info in mass insert: %+v", contact)
			continue
		}
		// The whole query will break if there are duplicates, so make sure there aren't any duplicates
		_, alreadyHandled := handledContacts[contact.JID]
		if alreadyHandled {
			s.log.Warnf("Duplicate contact info for %s in mass insert", contact.JID)
			continue
		}
		handledContacts[contact.JID] = struct{}{}
		baseIndex := i*3 + 1
		values = append(values, contact.JID.String(), contact.FirstName, contact.FullName)
		queryParts = append(queryParts, fmt.Sprintf(placeholderSyntax, baseIndex+1, baseIndex+2, baseIndex+3))
		i++
	}
	_, err := tx.Exec(fmt.Sprintf(putManyContactNamesQuery, strings.Join(queryParts, ",")), values...)
	return err
}

func (s *SQLStore) PutAllContactNames(contacts []store.ContactEntry) error {
	if len(contacts) > contactBatchSize {
		tx, err := s.db.Begin()
		if err != nil {
			return fmt.Errorf("failed to start transaction: %w", err)
		}
		for i := 0; i < len(contacts); i += contactBatchSize {
			var contactSlice []store.ContactEntry
			if len(contacts) > i+contactBatchSize {
				contactSlice = contacts[i : i+contactBatchSize]
			} else {
				contactSlice = contacts[i:]
			}
			err = s.putContactNamesBatch(tx, contactSlice)
			if err != nil {
				_ = tx.Rollback()
				return err
			}
		}
		err = tx.Commit()
		if err != nil {
			return fmt.Errorf("failed to commit transaction: %w", err)
		}
	} else if len(contacts) > 0 {
		err := s.putContactNamesBatch(s.db, contacts)
		if err != nil {
			return err
		}
	} else {
		return nil
	}
	s.contactCacheLock.Lock()
	// Just clear the cache, fetching pushnames and business names would be too much effort
	s.contactCache = make(map[types.JID]*types.ContactInfo)
	s.contactCacheLock.Unlock()
	return nil
}

func (s *SQLStore) getContact(user types.JID) (*types.ContactInfo, error) {
	cached, ok := s.contactCache[user]
	if ok {
		return cached, nil
	}

	var first, full, push, business sql.NullString
	err := s.db.QueryRow(getContactQuery, s.JID, user).Scan(&first, &full, &push, &business)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	info := &types.ContactInfo{
		Found:        err == nil,
		FirstName:    first.String,
		FullName:     full.String,
		PushName:     push.String,
		BusinessName: business.String,
	}
	s.contactCache[user] = info
	return info, nil
}

func (s *SQLStore) GetContact(user types.JID) (types.ContactInfo, error) {
	s.contactCacheLock.Lock()
	info, err := s.getContact(user)
	s.contactCacheLock.Unlock()
	if err != nil {
		return types.ContactInfo{}, err
	}
	return *info, nil
}

func (s *SQLStore) GetAllContacts() (map[types.JID]types.ContactInfo, error) {
	s.contactCacheLock.Lock()
	defer s.contactCacheLock.Unlock()
	rows, err := s.db.Query(getAllContactsQuery, s.JID)
	if err != nil {
		return nil, err
	}
	output := make(map[types.JID]types.ContactInfo, len(s.contactCache))
	for rows.Next() {
		var jid types.JID
		var first, full, push, business sql.NullString
		err = rows.Scan(&jid, &first, &full, &push, &business)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		info := types.ContactInfo{
			Found:        true,
			FirstName:    first.String,
			FullName:     full.String,
			PushName:     push.String,
			BusinessName: business.String,
		}
		output[jid] = info
		s.contactCache[jid] = &info
	}
	return output, nil
}

const (
	putChatSettingQuery = `INSERT INTO whatsmeow_chat_settings (our_jid, chat_jid, column_name)
	VALUES (?, ?, ?)
	ON DUPLICATE KEY UPDATE
		column_name = CASE
			WHEN VALUES(column_name) IS NOT NULL THEN VALUES(column_name)
			ELSE column_name
		END`
	getChatSettingsQuery = `
		SELECT muted_until, pinned, archived FROM whatsmeow_chat_settings WHERE our_jid=? AND chat_jid=?
	`
)

func (s *SQLStore) PutMutedUntil(chat types.JID, mutedUntil time.Time) error {
	var val int64
	if !mutedUntil.IsZero() {
		val = mutedUntil.Unix()
	}
	_, err := s.db.Exec(fmt.Sprintf(putChatSettingQuery, "muted_until"), s.JID, chat, val)
	return err
}

func (s *SQLStore) PutPinned(chat types.JID, pinned bool) error {
	_, err := s.db.Exec(fmt.Sprintf(putChatSettingQuery, "pinned"), s.JID, chat, pinned)
	return err
}

func (s *SQLStore) PutArchived(chat types.JID, archived bool) error {
	_, err := s.db.Exec(fmt.Sprintf(putChatSettingQuery, "archived"), s.JID, chat, archived)
	return err
}

func (s *SQLStore) GetChatSettings(chat types.JID) (settings types.LocalChatSettings, err error) {
	var mutedUntil int64
	err = s.db.QueryRow(getChatSettingsQuery, s.JID, chat).Scan(&mutedUntil, &settings.Pinned, &settings.Archived)
	if errors.Is(err, sql.ErrNoRows) {
		err = nil
	} else if err != nil {
		return
	} else {
		settings.Found = true
	}
	if mutedUntil != 0 {
		settings.MutedUntil = time.Unix(mutedUntil, 0)
	}
	return
}

const (
	putMsgSecret = `INSERT IGNORE INTO whatsmeow_message_secrets (our_jid, chat_jid, sender_jid, message_id, key)
	VALUES (?, ?, ?, ?, ?)`
	getMsgSecret = `
		SELECT ` + "`key`" + ` FROM whatsmeow_message_secrets WHERE our_jid=? AND chat_jid=? AND sender_jid=? AND message_id=?
	`
)

func (s *SQLStore) PutMessageSecrets(inserts []store.MessageSecretInsert) (err error) {
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	for _, insert := range inserts {
		_, err = tx.Exec(putMsgSecret, s.JID, insert.Chat.ToNonAD(), insert.Sender.ToNonAD(), insert.ID, insert.Secret)
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return
}

func (s *SQLStore) PutMessageSecret(chat, sender types.JID, id types.MessageID, secret []byte) (err error) {
	_, err = s.db.Exec(putMsgSecret, s.JID, chat.ToNonAD(), sender.ToNonAD(), id, secret)
	return
}

func (s *SQLStore) GetMessageSecret(chat, sender types.JID, id types.MessageID) (secret []byte, err error) {
	err = s.db.QueryRow(getMsgSecret, s.JID, chat.ToNonAD(), sender.ToNonAD(), id).Scan(&secret)
	if errors.Is(err, sql.ErrNoRows) {
		err = nil
	}
	return
}

const (
	putPrivacyTokens = `INSERT INTO whatsmeow_privacy_tokens (our_jid, their_jid, token, timestamp)
	VALUES (?, ?, ?, ?)
	ON DUPLICATE KEY UPDATE
		token = VALUES(token),
		timestamp = VALUES(timestamp)`

	getPrivacyToken = `SELECT token, timestamp FROM whatsmeow_privacy_tokens WHERE our_jid=? AND their_jid=?`
)

func (s *SQLStore) PutPrivacyTokens(tokens ...store.PrivacyToken) error {
	args := make([]any, 1+len(tokens)*3)
	placeholders := make([]string, len(tokens))
	args[0] = s.JID
	for i, token := range tokens {
		args[i*3+1] = token.User.ToNonAD().String()
		args[i*3+2] = token.Token
		args[i*3+3] = token.Timestamp.Unix()
		placeholders[i] = fmt.Sprintf("(?, $%d, $%d, $%d)", i*3+2, i*3+3, i*3+4)
	}
	query := strings.ReplaceAll(putPrivacyTokens, "(?, ?, ?, ?)", strings.Join(placeholders, ","))
	_, err := s.db.Exec(query, args...)
	return err
}

func (s *SQLStore) GetPrivacyToken(user types.JID) (*store.PrivacyToken, error) {
	var token store.PrivacyToken
	token.User = user.ToNonAD()
	var ts int64
	err := s.db.QueryRow(getPrivacyToken, s.JID, token.User).Scan(&token.Token, &ts)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	} else {
		token.Timestamp = time.Unix(ts, 0)
		return &token, nil
	}
}
