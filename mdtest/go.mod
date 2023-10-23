module github.com/kiriminaja/kaj-notification-srvc/src/pkg/whatsapp/whatsmeow/mdtest

go 1.19

require (
	github.com/mattn/go-sqlite3 v1.14.17
	github.com/mdp/qrterminal/v3 v3.0.0
	github.com/kiriminaja/kaj-notification-srvc/src/pkg/whatsapp/whatsmeow v0.0.0-20230805111647-405414b9b5c0
	google.golang.org/protobuf v1.31.0
)

require (
	filippo.io/edwards25519 v1.0.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	go.mau.fi/libsignal v0.1.0 // indirect
	go.mau.fi/util v0.2.0 // indirect
	golang.org/x/crypto v0.14.0 // indirect
	rsc.io/qr v0.2.0 // indirect
)

replace github.com/kiriminaja/kaj-notification-srvc/src/pkg/whatsapp/whatsmeow => ../
