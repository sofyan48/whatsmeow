// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: waArmadilloXMA/WAArmadilloXMA.proto

package waArmadilloXMA

import (
	waCommon "github.com/sofyan48/whatsmeow/binary/armadillo/waCommon"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

import _ "embed"

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ExtendedContentMessage_OverlayIconGlyph int32

const (
	ExtendedContentMessage_INFO            ExtendedContentMessage_OverlayIconGlyph = 0
	ExtendedContentMessage_EYE_OFF         ExtendedContentMessage_OverlayIconGlyph = 1
	ExtendedContentMessage_NEWS_OFF        ExtendedContentMessage_OverlayIconGlyph = 2
	ExtendedContentMessage_WARNING         ExtendedContentMessage_OverlayIconGlyph = 3
	ExtendedContentMessage_PRIVATE         ExtendedContentMessage_OverlayIconGlyph = 4
	ExtendedContentMessage_NONE            ExtendedContentMessage_OverlayIconGlyph = 5
	ExtendedContentMessage_MEDIA_LABEL     ExtendedContentMessage_OverlayIconGlyph = 6
	ExtendedContentMessage_POST_COVER      ExtendedContentMessage_OverlayIconGlyph = 7
	ExtendedContentMessage_POST_LABEL      ExtendedContentMessage_OverlayIconGlyph = 8
	ExtendedContentMessage_WARNING_SCREENS ExtendedContentMessage_OverlayIconGlyph = 9
)

// Enum value maps for ExtendedContentMessage_OverlayIconGlyph.
var (
	ExtendedContentMessage_OverlayIconGlyph_name = map[int32]string{
		0: "INFO",
		1: "EYE_OFF",
		2: "NEWS_OFF",
		3: "WARNING",
		4: "PRIVATE",
		5: "NONE",
		6: "MEDIA_LABEL",
		7: "POST_COVER",
		8: "POST_LABEL",
		9: "WARNING_SCREENS",
	}
	ExtendedContentMessage_OverlayIconGlyph_value = map[string]int32{
		"INFO":            0,
		"EYE_OFF":         1,
		"NEWS_OFF":        2,
		"WARNING":         3,
		"PRIVATE":         4,
		"NONE":            5,
		"MEDIA_LABEL":     6,
		"POST_COVER":      7,
		"POST_LABEL":      8,
		"WARNING_SCREENS": 9,
	}
)

func (x ExtendedContentMessage_OverlayIconGlyph) Enum() *ExtendedContentMessage_OverlayIconGlyph {
	p := new(ExtendedContentMessage_OverlayIconGlyph)
	*p = x
	return p
}

func (x ExtendedContentMessage_OverlayIconGlyph) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExtendedContentMessage_OverlayIconGlyph) Descriptor() protoreflect.EnumDescriptor {
	return file_waArmadilloXMA_WAArmadilloXMA_proto_enumTypes[0].Descriptor()
}

func (ExtendedContentMessage_OverlayIconGlyph) Type() protoreflect.EnumType {
	return &file_waArmadilloXMA_WAArmadilloXMA_proto_enumTypes[0]
}

func (x ExtendedContentMessage_OverlayIconGlyph) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExtendedContentMessage_OverlayIconGlyph.Descriptor instead.
func (ExtendedContentMessage_OverlayIconGlyph) EnumDescriptor() ([]byte, []int) {
	return file_waArmadilloXMA_WAArmadilloXMA_proto_rawDescGZIP(), []int{0, 0}
}

type ExtendedContentMessage_CtaButtonType int32

const (
	ExtendedContentMessage_CTABUTTONTYPE_UNKNOWN ExtendedContentMessage_CtaButtonType = 0
	ExtendedContentMessage_OPEN_NATIVE           ExtendedContentMessage_CtaButtonType = 11
)

// Enum value maps for ExtendedContentMessage_CtaButtonType.
var (
	ExtendedContentMessage_CtaButtonType_name = map[int32]string{
		0:  "CTABUTTONTYPE_UNKNOWN",
		11: "OPEN_NATIVE",
	}
	ExtendedContentMessage_CtaButtonType_value = map[string]int32{
		"CTABUTTONTYPE_UNKNOWN": 0,
		"OPEN_NATIVE":           11,
	}
)

func (x ExtendedContentMessage_CtaButtonType) Enum() *ExtendedContentMessage_CtaButtonType {
	p := new(ExtendedContentMessage_CtaButtonType)
	*p = x
	return p
}

func (x ExtendedContentMessage_CtaButtonType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExtendedContentMessage_CtaButtonType) Descriptor() protoreflect.EnumDescriptor {
	return file_waArmadilloXMA_WAArmadilloXMA_proto_enumTypes[1].Descriptor()
}

func (ExtendedContentMessage_CtaButtonType) Type() protoreflect.EnumType {
	return &file_waArmadilloXMA_WAArmadilloXMA_proto_enumTypes[1]
}

func (x ExtendedContentMessage_CtaButtonType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExtendedContentMessage_CtaButtonType.Descriptor instead.
func (ExtendedContentMessage_CtaButtonType) EnumDescriptor() ([]byte, []int) {
	return file_waArmadilloXMA_WAArmadilloXMA_proto_rawDescGZIP(), []int{0, 1}
}

type ExtendedContentMessage_XmaLayoutType int32

const (
	ExtendedContentMessage_SINGLE        ExtendedContentMessage_XmaLayoutType = 0
	ExtendedContentMessage_PORTRAIT      ExtendedContentMessage_XmaLayoutType = 3
	ExtendedContentMessage_STANDARD_DXMA ExtendedContentMessage_XmaLayoutType = 12
	ExtendedContentMessage_LIST_DXMA     ExtendedContentMessage_XmaLayoutType = 15
)

// Enum value maps for ExtendedContentMessage_XmaLayoutType.
var (
	ExtendedContentMessage_XmaLayoutType_name = map[int32]string{
		0:  "SINGLE",
		3:  "PORTRAIT",
		12: "STANDARD_DXMA",
		15: "LIST_DXMA",
	}
	ExtendedContentMessage_XmaLayoutType_value = map[string]int32{
		"SINGLE":        0,
		"PORTRAIT":      3,
		"STANDARD_DXMA": 12,
		"LIST_DXMA":     15,
	}
)

func (x ExtendedContentMessage_XmaLayoutType) Enum() *ExtendedContentMessage_XmaLayoutType {
	p := new(ExtendedContentMessage_XmaLayoutType)
	*p = x
	return p
}

func (x ExtendedContentMessage_XmaLayoutType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExtendedContentMessage_XmaLayoutType) Descriptor() protoreflect.EnumDescriptor {
	return file_waArmadilloXMA_WAArmadilloXMA_proto_enumTypes[2].Descriptor()
}

func (ExtendedContentMessage_XmaLayoutType) Type() protoreflect.EnumType {
	return &file_waArmadilloXMA_WAArmadilloXMA_proto_enumTypes[2]
}

func (x ExtendedContentMessage_XmaLayoutType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExtendedContentMessage_XmaLayoutType.Descriptor instead.
func (ExtendedContentMessage_XmaLayoutType) EnumDescriptor() ([]byte, []int) {
	return file_waArmadilloXMA_WAArmadilloXMA_proto_rawDescGZIP(), []int{0, 2}
}

type ExtendedContentMessage_ExtendedContentType int32

const (
	ExtendedContentMessage_EXTENDEDCONTENTTYPE_UNKNOWN             ExtendedContentMessage_ExtendedContentType = 0
	ExtendedContentMessage_IG_STORY_PHOTO_MENTION                  ExtendedContentMessage_ExtendedContentType = 4
	ExtendedContentMessage_IG_SINGLE_IMAGE_POST_SHARE              ExtendedContentMessage_ExtendedContentType = 9
	ExtendedContentMessage_IG_MULTIPOST_SHARE                      ExtendedContentMessage_ExtendedContentType = 10
	ExtendedContentMessage_IG_SINGLE_VIDEO_POST_SHARE              ExtendedContentMessage_ExtendedContentType = 11
	ExtendedContentMessage_IG_STORY_PHOTO_SHARE                    ExtendedContentMessage_ExtendedContentType = 12
	ExtendedContentMessage_IG_STORY_VIDEO_SHARE                    ExtendedContentMessage_ExtendedContentType = 13
	ExtendedContentMessage_IG_CLIPS_SHARE                          ExtendedContentMessage_ExtendedContentType = 14
	ExtendedContentMessage_IG_IGTV_SHARE                           ExtendedContentMessage_ExtendedContentType = 15
	ExtendedContentMessage_IG_SHOP_SHARE                           ExtendedContentMessage_ExtendedContentType = 16
	ExtendedContentMessage_IG_PROFILE_SHARE                        ExtendedContentMessage_ExtendedContentType = 19
	ExtendedContentMessage_IG_STORY_PHOTO_HIGHLIGHT_SHARE          ExtendedContentMessage_ExtendedContentType = 20
	ExtendedContentMessage_IG_STORY_VIDEO_HIGHLIGHT_SHARE          ExtendedContentMessage_ExtendedContentType = 21
	ExtendedContentMessage_IG_STORY_REPLY                          ExtendedContentMessage_ExtendedContentType = 22
	ExtendedContentMessage_IG_STORY_REACTION                       ExtendedContentMessage_ExtendedContentType = 23
	ExtendedContentMessage_IG_STORY_VIDEO_MENTION                  ExtendedContentMessage_ExtendedContentType = 24
	ExtendedContentMessage_IG_STORY_HIGHLIGHT_REPLY                ExtendedContentMessage_ExtendedContentType = 25
	ExtendedContentMessage_IG_STORY_HIGHLIGHT_REACTION             ExtendedContentMessage_ExtendedContentType = 26
	ExtendedContentMessage_IG_EXTERNAL_LINK                        ExtendedContentMessage_ExtendedContentType = 27
	ExtendedContentMessage_IG_RECEIVER_FETCH                       ExtendedContentMessage_ExtendedContentType = 28
	ExtendedContentMessage_FB_FEED_SHARE                           ExtendedContentMessage_ExtendedContentType = 1000
	ExtendedContentMessage_FB_STORY_REPLY                          ExtendedContentMessage_ExtendedContentType = 1001
	ExtendedContentMessage_FB_STORY_SHARE                          ExtendedContentMessage_ExtendedContentType = 1002
	ExtendedContentMessage_FB_STORY_MENTION                        ExtendedContentMessage_ExtendedContentType = 1003
	ExtendedContentMessage_FB_FEED_VIDEO_SHARE                     ExtendedContentMessage_ExtendedContentType = 1004
	ExtendedContentMessage_FB_GAMING_CUSTOM_UPDATE                 ExtendedContentMessage_ExtendedContentType = 1005
	ExtendedContentMessage_FB_PRODUCER_STORY_REPLY                 ExtendedContentMessage_ExtendedContentType = 1006
	ExtendedContentMessage_FB_EVENT                                ExtendedContentMessage_ExtendedContentType = 1007
	ExtendedContentMessage_FB_FEED_POST_PRIVATE_REPLY              ExtendedContentMessage_ExtendedContentType = 1008
	ExtendedContentMessage_FB_SHORT                                ExtendedContentMessage_ExtendedContentType = 1009
	ExtendedContentMessage_FB_COMMENT_MENTION_SHARE                ExtendedContentMessage_ExtendedContentType = 1010
	ExtendedContentMessage_MSG_EXTERNAL_LINK_SHARE                 ExtendedContentMessage_ExtendedContentType = 2000
	ExtendedContentMessage_MSG_P2P_PAYMENT                         ExtendedContentMessage_ExtendedContentType = 2001
	ExtendedContentMessage_MSG_LOCATION_SHARING                    ExtendedContentMessage_ExtendedContentType = 2002
	ExtendedContentMessage_MSG_LOCATION_SHARING_V2                 ExtendedContentMessage_ExtendedContentType = 2003
	ExtendedContentMessage_MSG_HIGHLIGHTS_TAB_FRIEND_UPDATES_REPLY ExtendedContentMessage_ExtendedContentType = 2004
	ExtendedContentMessage_MSG_HIGHLIGHTS_TAB_LOCAL_EVENT_REPLY    ExtendedContentMessage_ExtendedContentType = 2005
	ExtendedContentMessage_MSG_RECEIVER_FETCH                      ExtendedContentMessage_ExtendedContentType = 2006
	ExtendedContentMessage_MSG_IG_MEDIA_SHARE                      ExtendedContentMessage_ExtendedContentType = 2007
	ExtendedContentMessage_MSG_GEN_AI_SEARCH_PLUGIN_RESPONSE       ExtendedContentMessage_ExtendedContentType = 2008
	ExtendedContentMessage_MSG_REELS_LIST                          ExtendedContentMessage_ExtendedContentType = 2009
	ExtendedContentMessage_MSG_CONTACT                             ExtendedContentMessage_ExtendedContentType = 2010
	ExtendedContentMessage_RTC_AUDIO_CALL                          ExtendedContentMessage_ExtendedContentType = 3000
	ExtendedContentMessage_RTC_VIDEO_CALL                          ExtendedContentMessage_ExtendedContentType = 3001
	ExtendedContentMessage_RTC_MISSED_AUDIO_CALL                   ExtendedContentMessage_ExtendedContentType = 3002
	ExtendedContentMessage_RTC_MISSED_VIDEO_CALL                   ExtendedContentMessage_ExtendedContentType = 3003
	ExtendedContentMessage_RTC_GROUP_AUDIO_CALL                    ExtendedContentMessage_ExtendedContentType = 3004
	ExtendedContentMessage_RTC_GROUP_VIDEO_CALL                    ExtendedContentMessage_ExtendedContentType = 3005
	ExtendedContentMessage_RTC_MISSED_GROUP_AUDIO_CALL             ExtendedContentMessage_ExtendedContentType = 3006
	ExtendedContentMessage_RTC_MISSED_GROUP_VIDEO_CALL             ExtendedContentMessage_ExtendedContentType = 3007
	ExtendedContentMessage_DATACLASS_SENDER_COPY                   ExtendedContentMessage_ExtendedContentType = 4000
)

// Enum value maps for ExtendedContentMessage_ExtendedContentType.
var (
	ExtendedContentMessage_ExtendedContentType_name = map[int32]string{
		0:    "EXTENDEDCONTENTTYPE_UNKNOWN",
		4:    "IG_STORY_PHOTO_MENTION",
		9:    "IG_SINGLE_IMAGE_POST_SHARE",
		10:   "IG_MULTIPOST_SHARE",
		11:   "IG_SINGLE_VIDEO_POST_SHARE",
		12:   "IG_STORY_PHOTO_SHARE",
		13:   "IG_STORY_VIDEO_SHARE",
		14:   "IG_CLIPS_SHARE",
		15:   "IG_IGTV_SHARE",
		16:   "IG_SHOP_SHARE",
		19:   "IG_PROFILE_SHARE",
		20:   "IG_STORY_PHOTO_HIGHLIGHT_SHARE",
		21:   "IG_STORY_VIDEO_HIGHLIGHT_SHARE",
		22:   "IG_STORY_REPLY",
		23:   "IG_STORY_REACTION",
		24:   "IG_STORY_VIDEO_MENTION",
		25:   "IG_STORY_HIGHLIGHT_REPLY",
		26:   "IG_STORY_HIGHLIGHT_REACTION",
		27:   "IG_EXTERNAL_LINK",
		28:   "IG_RECEIVER_FETCH",
		1000: "FB_FEED_SHARE",
		1001: "FB_STORY_REPLY",
		1002: "FB_STORY_SHARE",
		1003: "FB_STORY_MENTION",
		1004: "FB_FEED_VIDEO_SHARE",
		1005: "FB_GAMING_CUSTOM_UPDATE",
		1006: "FB_PRODUCER_STORY_REPLY",
		1007: "FB_EVENT",
		1008: "FB_FEED_POST_PRIVATE_REPLY",
		1009: "FB_SHORT",
		1010: "FB_COMMENT_MENTION_SHARE",
		2000: "MSG_EXTERNAL_LINK_SHARE",
		2001: "MSG_P2P_PAYMENT",
		2002: "MSG_LOCATION_SHARING",
		2003: "MSG_LOCATION_SHARING_V2",
		2004: "MSG_HIGHLIGHTS_TAB_FRIEND_UPDATES_REPLY",
		2005: "MSG_HIGHLIGHTS_TAB_LOCAL_EVENT_REPLY",
		2006: "MSG_RECEIVER_FETCH",
		2007: "MSG_IG_MEDIA_SHARE",
		2008: "MSG_GEN_AI_SEARCH_PLUGIN_RESPONSE",
		2009: "MSG_REELS_LIST",
		2010: "MSG_CONTACT",
		3000: "RTC_AUDIO_CALL",
		3001: "RTC_VIDEO_CALL",
		3002: "RTC_MISSED_AUDIO_CALL",
		3003: "RTC_MISSED_VIDEO_CALL",
		3004: "RTC_GROUP_AUDIO_CALL",
		3005: "RTC_GROUP_VIDEO_CALL",
		3006: "RTC_MISSED_GROUP_AUDIO_CALL",
		3007: "RTC_MISSED_GROUP_VIDEO_CALL",
		4000: "DATACLASS_SENDER_COPY",
	}
	ExtendedContentMessage_ExtendedContentType_value = map[string]int32{
		"EXTENDEDCONTENTTYPE_UNKNOWN":             0,
		"IG_STORY_PHOTO_MENTION":                  4,
		"IG_SINGLE_IMAGE_POST_SHARE":              9,
		"IG_MULTIPOST_SHARE":                      10,
		"IG_SINGLE_VIDEO_POST_SHARE":              11,
		"IG_STORY_PHOTO_SHARE":                    12,
		"IG_STORY_VIDEO_SHARE":                    13,
		"IG_CLIPS_SHARE":                          14,
		"IG_IGTV_SHARE":                           15,
		"IG_SHOP_SHARE":                           16,
		"IG_PROFILE_SHARE":                        19,
		"IG_STORY_PHOTO_HIGHLIGHT_SHARE":          20,
		"IG_STORY_VIDEO_HIGHLIGHT_SHARE":          21,
		"IG_STORY_REPLY":                          22,
		"IG_STORY_REACTION":                       23,
		"IG_STORY_VIDEO_MENTION":                  24,
		"IG_STORY_HIGHLIGHT_REPLY":                25,
		"IG_STORY_HIGHLIGHT_REACTION":             26,
		"IG_EXTERNAL_LINK":                        27,
		"IG_RECEIVER_FETCH":                       28,
		"FB_FEED_SHARE":                           1000,
		"FB_STORY_REPLY":                          1001,
		"FB_STORY_SHARE":                          1002,
		"FB_STORY_MENTION":                        1003,
		"FB_FEED_VIDEO_SHARE":                     1004,
		"FB_GAMING_CUSTOM_UPDATE":                 1005,
		"FB_PRODUCER_STORY_REPLY":                 1006,
		"FB_EVENT":                                1007,
		"FB_FEED_POST_PRIVATE_REPLY":              1008,
		"FB_SHORT":                                1009,
		"FB_COMMENT_MENTION_SHARE":                1010,
		"MSG_EXTERNAL_LINK_SHARE":                 2000,
		"MSG_P2P_PAYMENT":                         2001,
		"MSG_LOCATION_SHARING":                    2002,
		"MSG_LOCATION_SHARING_V2":                 2003,
		"MSG_HIGHLIGHTS_TAB_FRIEND_UPDATES_REPLY": 2004,
		"MSG_HIGHLIGHTS_TAB_LOCAL_EVENT_REPLY":    2005,
		"MSG_RECEIVER_FETCH":                      2006,
		"MSG_IG_MEDIA_SHARE":                      2007,
		"MSG_GEN_AI_SEARCH_PLUGIN_RESPONSE":       2008,
		"MSG_REELS_LIST":                          2009,
		"MSG_CONTACT":                             2010,
		"RTC_AUDIO_CALL":                          3000,
		"RTC_VIDEO_CALL":                          3001,
		"RTC_MISSED_AUDIO_CALL":                   3002,
		"RTC_MISSED_VIDEO_CALL":                   3003,
		"RTC_GROUP_AUDIO_CALL":                    3004,
		"RTC_GROUP_VIDEO_CALL":                    3005,
		"RTC_MISSED_GROUP_AUDIO_CALL":             3006,
		"RTC_MISSED_GROUP_VIDEO_CALL":             3007,
		"DATACLASS_SENDER_COPY":                   4000,
	}
)

func (x ExtendedContentMessage_ExtendedContentType) Enum() *ExtendedContentMessage_ExtendedContentType {
	p := new(ExtendedContentMessage_ExtendedContentType)
	*p = x
	return p
}

func (x ExtendedContentMessage_ExtendedContentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExtendedContentMessage_ExtendedContentType) Descriptor() protoreflect.EnumDescriptor {
	return file_waArmadilloXMA_WAArmadilloXMA_proto_enumTypes[3].Descriptor()
}

func (ExtendedContentMessage_ExtendedContentType) Type() protoreflect.EnumType {
	return &file_waArmadilloXMA_WAArmadilloXMA_proto_enumTypes[3]
}

func (x ExtendedContentMessage_ExtendedContentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExtendedContentMessage_ExtendedContentType.Descriptor instead.
func (ExtendedContentMessage_ExtendedContentType) EnumDescriptor() ([]byte, []int) {
	return file_waArmadilloXMA_WAArmadilloXMA_proto_rawDescGZIP(), []int{0, 3}
}

type ExtendedContentMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssociatedMessage     *waCommon.SubProtocol                      `protobuf:"bytes,1,opt,name=associatedMessage,proto3" json:"associatedMessage,omitempty"`
	TargetType            ExtendedContentMessage_ExtendedContentType `protobuf:"varint,2,opt,name=targetType,proto3,enum=WAArmadilloXMA.ExtendedContentMessage_ExtendedContentType" json:"targetType,omitempty"`
	TargetUsername        string                                     `protobuf:"bytes,3,opt,name=targetUsername,proto3" json:"targetUsername,omitempty"`
	TargetID              string                                     `protobuf:"bytes,4,opt,name=targetID,proto3" json:"targetID,omitempty"`
	TargetExpiringAtSec   int64                                      `protobuf:"varint,5,opt,name=targetExpiringAtSec,proto3" json:"targetExpiringAtSec,omitempty"`
	XmaLayoutType         ExtendedContentMessage_XmaLayoutType       `protobuf:"varint,6,opt,name=xmaLayoutType,proto3,enum=WAArmadilloXMA.ExtendedContentMessage_XmaLayoutType" json:"xmaLayoutType,omitempty"`
	Ctas                  []*ExtendedContentMessage_CTA              `protobuf:"bytes,7,rep,name=ctas,proto3" json:"ctas,omitempty"`
	Previews              []*waCommon.SubProtocol                    `protobuf:"bytes,8,rep,name=previews,proto3" json:"previews,omitempty"`
	TitleText             string                                     `protobuf:"bytes,9,opt,name=titleText,proto3" json:"titleText,omitempty"`
	SubtitleText          string                                     `protobuf:"bytes,10,opt,name=subtitleText,proto3" json:"subtitleText,omitempty"`
	MaxTitleNumOfLines    uint32                                     `protobuf:"varint,11,opt,name=maxTitleNumOfLines,proto3" json:"maxTitleNumOfLines,omitempty"`
	MaxSubtitleNumOfLines uint32                                     `protobuf:"varint,12,opt,name=maxSubtitleNumOfLines,proto3" json:"maxSubtitleNumOfLines,omitempty"`
	Favicon               *waCommon.SubProtocol                      `protobuf:"bytes,13,opt,name=favicon,proto3" json:"favicon,omitempty"`
	HeaderImage           *waCommon.SubProtocol                      `protobuf:"bytes,14,opt,name=headerImage,proto3" json:"headerImage,omitempty"`
	HeaderTitle           string                                     `protobuf:"bytes,15,opt,name=headerTitle,proto3" json:"headerTitle,omitempty"`
	OverlayIconGlyph      ExtendedContentMessage_OverlayIconGlyph    `protobuf:"varint,16,opt,name=overlayIconGlyph,proto3,enum=WAArmadilloXMA.ExtendedContentMessage_OverlayIconGlyph" json:"overlayIconGlyph,omitempty"`
	OverlayTitle          string                                     `protobuf:"bytes,17,opt,name=overlayTitle,proto3" json:"overlayTitle,omitempty"`
	OverlayDescription    string                                     `protobuf:"bytes,18,opt,name=overlayDescription,proto3" json:"overlayDescription,omitempty"`
	SentWithMessageID     string                                     `protobuf:"bytes,19,opt,name=sentWithMessageID,proto3" json:"sentWithMessageID,omitempty"`
	MessageText           string                                     `protobuf:"bytes,20,opt,name=messageText,proto3" json:"messageText,omitempty"`
	HeaderSubtitle        string                                     `protobuf:"bytes,21,opt,name=headerSubtitle,proto3" json:"headerSubtitle,omitempty"`
	XmaDataclass          string                                     `protobuf:"bytes,22,opt,name=xmaDataclass,proto3" json:"xmaDataclass,omitempty"`
	ContentRef            string                                     `protobuf:"bytes,23,opt,name=contentRef,proto3" json:"contentRef,omitempty"`
}

func (x *ExtendedContentMessage) Reset() {
	*x = ExtendedContentMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waArmadilloXMA_WAArmadilloXMA_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtendedContentMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtendedContentMessage) ProtoMessage() {}

func (x *ExtendedContentMessage) ProtoReflect() protoreflect.Message {
	mi := &file_waArmadilloXMA_WAArmadilloXMA_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtendedContentMessage.ProtoReflect.Descriptor instead.
func (*ExtendedContentMessage) Descriptor() ([]byte, []int) {
	return file_waArmadilloXMA_WAArmadilloXMA_proto_rawDescGZIP(), []int{0}
}

func (x *ExtendedContentMessage) GetAssociatedMessage() *waCommon.SubProtocol {
	if x != nil {
		return x.AssociatedMessage
	}
	return nil
}

func (x *ExtendedContentMessage) GetTargetType() ExtendedContentMessage_ExtendedContentType {
	if x != nil {
		return x.TargetType
	}
	return ExtendedContentMessage_EXTENDEDCONTENTTYPE_UNKNOWN
}

func (x *ExtendedContentMessage) GetTargetUsername() string {
	if x != nil {
		return x.TargetUsername
	}
	return ""
}

func (x *ExtendedContentMessage) GetTargetID() string {
	if x != nil {
		return x.TargetID
	}
	return ""
}

func (x *ExtendedContentMessage) GetTargetExpiringAtSec() int64 {
	if x != nil {
		return x.TargetExpiringAtSec
	}
	return 0
}

func (x *ExtendedContentMessage) GetXmaLayoutType() ExtendedContentMessage_XmaLayoutType {
	if x != nil {
		return x.XmaLayoutType
	}
	return ExtendedContentMessage_SINGLE
}

func (x *ExtendedContentMessage) GetCtas() []*ExtendedContentMessage_CTA {
	if x != nil {
		return x.Ctas
	}
	return nil
}

func (x *ExtendedContentMessage) GetPreviews() []*waCommon.SubProtocol {
	if x != nil {
		return x.Previews
	}
	return nil
}

func (x *ExtendedContentMessage) GetTitleText() string {
	if x != nil {
		return x.TitleText
	}
	return ""
}

func (x *ExtendedContentMessage) GetSubtitleText() string {
	if x != nil {
		return x.SubtitleText
	}
	return ""
}

func (x *ExtendedContentMessage) GetMaxTitleNumOfLines() uint32 {
	if x != nil {
		return x.MaxTitleNumOfLines
	}
	return 0
}

func (x *ExtendedContentMessage) GetMaxSubtitleNumOfLines() uint32 {
	if x != nil {
		return x.MaxSubtitleNumOfLines
	}
	return 0
}

func (x *ExtendedContentMessage) GetFavicon() *waCommon.SubProtocol {
	if x != nil {
		return x.Favicon
	}
	return nil
}

func (x *ExtendedContentMessage) GetHeaderImage() *waCommon.SubProtocol {
	if x != nil {
		return x.HeaderImage
	}
	return nil
}

func (x *ExtendedContentMessage) GetHeaderTitle() string {
	if x != nil {
		return x.HeaderTitle
	}
	return ""
}

func (x *ExtendedContentMessage) GetOverlayIconGlyph() ExtendedContentMessage_OverlayIconGlyph {
	if x != nil {
		return x.OverlayIconGlyph
	}
	return ExtendedContentMessage_INFO
}

func (x *ExtendedContentMessage) GetOverlayTitle() string {
	if x != nil {
		return x.OverlayTitle
	}
	return ""
}

func (x *ExtendedContentMessage) GetOverlayDescription() string {
	if x != nil {
		return x.OverlayDescription
	}
	return ""
}

func (x *ExtendedContentMessage) GetSentWithMessageID() string {
	if x != nil {
		return x.SentWithMessageID
	}
	return ""
}

func (x *ExtendedContentMessage) GetMessageText() string {
	if x != nil {
		return x.MessageText
	}
	return ""
}

func (x *ExtendedContentMessage) GetHeaderSubtitle() string {
	if x != nil {
		return x.HeaderSubtitle
	}
	return ""
}

func (x *ExtendedContentMessage) GetXmaDataclass() string {
	if x != nil {
		return x.XmaDataclass
	}
	return ""
}

func (x *ExtendedContentMessage) GetContentRef() string {
	if x != nil {
		return x.ContentRef
	}
	return ""
}

type ExtendedContentMessage_CTA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ButtonType ExtendedContentMessage_CtaButtonType `protobuf:"varint,1,opt,name=buttonType,proto3,enum=WAArmadilloXMA.ExtendedContentMessage_CtaButtonType" json:"buttonType,omitempty"`
	Title      string                               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	ActionURL  string                               `protobuf:"bytes,3,opt,name=actionURL,proto3" json:"actionURL,omitempty"`
	NativeURL  string                               `protobuf:"bytes,4,opt,name=nativeURL,proto3" json:"nativeURL,omitempty"`
	CtaType    string                               `protobuf:"bytes,5,opt,name=ctaType,proto3" json:"ctaType,omitempty"`
}

func (x *ExtendedContentMessage_CTA) Reset() {
	*x = ExtendedContentMessage_CTA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waArmadilloXMA_WAArmadilloXMA_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtendedContentMessage_CTA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtendedContentMessage_CTA) ProtoMessage() {}

func (x *ExtendedContentMessage_CTA) ProtoReflect() protoreflect.Message {
	mi := &file_waArmadilloXMA_WAArmadilloXMA_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtendedContentMessage_CTA.ProtoReflect.Descriptor instead.
func (*ExtendedContentMessage_CTA) Descriptor() ([]byte, []int) {
	return file_waArmadilloXMA_WAArmadilloXMA_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ExtendedContentMessage_CTA) GetButtonType() ExtendedContentMessage_CtaButtonType {
	if x != nil {
		return x.ButtonType
	}
	return ExtendedContentMessage_CTABUTTONTYPE_UNKNOWN
}

func (x *ExtendedContentMessage_CTA) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ExtendedContentMessage_CTA) GetActionURL() string {
	if x != nil {
		return x.ActionURL
	}
	return ""
}

func (x *ExtendedContentMessage_CTA) GetNativeURL() string {
	if x != nil {
		return x.NativeURL
	}
	return ""
}

func (x *ExtendedContentMessage_CTA) GetCtaType() string {
	if x != nil {
		return x.CtaType
	}
	return ""
}

var File_waArmadilloXMA_WAArmadilloXMA_proto protoreflect.FileDescriptor

//go:embed WAArmadilloXMA.pb.raw
var file_waArmadilloXMA_WAArmadilloXMA_proto_rawDesc []byte

var (
	file_waArmadilloXMA_WAArmadilloXMA_proto_rawDescOnce sync.Once
	file_waArmadilloXMA_WAArmadilloXMA_proto_rawDescData = file_waArmadilloXMA_WAArmadilloXMA_proto_rawDesc
)

func file_waArmadilloXMA_WAArmadilloXMA_proto_rawDescGZIP() []byte {
	file_waArmadilloXMA_WAArmadilloXMA_proto_rawDescOnce.Do(func() {
		file_waArmadilloXMA_WAArmadilloXMA_proto_rawDescData = protoimpl.X.CompressGZIP(file_waArmadilloXMA_WAArmadilloXMA_proto_rawDescData)
	})
	return file_waArmadilloXMA_WAArmadilloXMA_proto_rawDescData
}

var file_waArmadilloXMA_WAArmadilloXMA_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_waArmadilloXMA_WAArmadilloXMA_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_waArmadilloXMA_WAArmadilloXMA_proto_goTypes = []interface{}{
	(ExtendedContentMessage_OverlayIconGlyph)(0),    // 0: WAArmadilloXMA.ExtendedContentMessage.OverlayIconGlyph
	(ExtendedContentMessage_CtaButtonType)(0),       // 1: WAArmadilloXMA.ExtendedContentMessage.CtaButtonType
	(ExtendedContentMessage_XmaLayoutType)(0),       // 2: WAArmadilloXMA.ExtendedContentMessage.XmaLayoutType
	(ExtendedContentMessage_ExtendedContentType)(0), // 3: WAArmadilloXMA.ExtendedContentMessage.ExtendedContentType
	(*ExtendedContentMessage)(nil),                  // 4: WAArmadilloXMA.ExtendedContentMessage
	(*ExtendedContentMessage_CTA)(nil),              // 5: WAArmadilloXMA.ExtendedContentMessage.CTA
	(*waCommon.SubProtocol)(nil),                    // 6: WACommon.SubProtocol
}
var file_waArmadilloXMA_WAArmadilloXMA_proto_depIdxs = []int32{
	6, // 0: WAArmadilloXMA.ExtendedContentMessage.associatedMessage:type_name -> WACommon.SubProtocol
	3, // 1: WAArmadilloXMA.ExtendedContentMessage.targetType:type_name -> WAArmadilloXMA.ExtendedContentMessage.ExtendedContentType
	2, // 2: WAArmadilloXMA.ExtendedContentMessage.xmaLayoutType:type_name -> WAArmadilloXMA.ExtendedContentMessage.XmaLayoutType
	5, // 3: WAArmadilloXMA.ExtendedContentMessage.ctas:type_name -> WAArmadilloXMA.ExtendedContentMessage.CTA
	6, // 4: WAArmadilloXMA.ExtendedContentMessage.previews:type_name -> WACommon.SubProtocol
	6, // 5: WAArmadilloXMA.ExtendedContentMessage.favicon:type_name -> WACommon.SubProtocol
	6, // 6: WAArmadilloXMA.ExtendedContentMessage.headerImage:type_name -> WACommon.SubProtocol
	0, // 7: WAArmadilloXMA.ExtendedContentMessage.overlayIconGlyph:type_name -> WAArmadilloXMA.ExtendedContentMessage.OverlayIconGlyph
	1, // 8: WAArmadilloXMA.ExtendedContentMessage.CTA.buttonType:type_name -> WAArmadilloXMA.ExtendedContentMessage.CtaButtonType
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_waArmadilloXMA_WAArmadilloXMA_proto_init() }
func file_waArmadilloXMA_WAArmadilloXMA_proto_init() {
	if File_waArmadilloXMA_WAArmadilloXMA_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_waArmadilloXMA_WAArmadilloXMA_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtendedContentMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_waArmadilloXMA_WAArmadilloXMA_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtendedContentMessage_CTA); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_waArmadilloXMA_WAArmadilloXMA_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_waArmadilloXMA_WAArmadilloXMA_proto_goTypes,
		DependencyIndexes: file_waArmadilloXMA_WAArmadilloXMA_proto_depIdxs,
		EnumInfos:         file_waArmadilloXMA_WAArmadilloXMA_proto_enumTypes,
		MessageInfos:      file_waArmadilloXMA_WAArmadilloXMA_proto_msgTypes,
	}.Build()
	File_waArmadilloXMA_WAArmadilloXMA_proto = out.File
	file_waArmadilloXMA_WAArmadilloXMA_proto_rawDesc = nil
	file_waArmadilloXMA_WAArmadilloXMA_proto_goTypes = nil
	file_waArmadilloXMA_WAArmadilloXMA_proto_depIdxs = nil
}
