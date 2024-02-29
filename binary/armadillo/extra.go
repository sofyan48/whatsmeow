package armadillo

import (
	"github.com/sofyan48/whatsmeow/binary/armadillo/waArmadilloApplication"
	"github.com/sofyan48/whatsmeow/binary/armadillo/waCommon"
	"github.com/sofyan48/whatsmeow/binary/armadillo/waConsumerApplication"
	"github.com/sofyan48/whatsmeow/binary/armadillo/waMultiDevice"
)

type MessageApplicationSub interface {
	IsMessageApplicationSub()
}

type Unsupported_BusinessApplication waCommon.SubProtocol
type Unsupported_PaymentApplication waCommon.SubProtocol
type Unsupported_Voip waCommon.SubProtocol

var (
	_ MessageApplicationSub = (*waConsumerApplication.ConsumerApplication)(nil) // 2
	_ MessageApplicationSub = (*Unsupported_BusinessApplication)(nil)           // 3
	_ MessageApplicationSub = (*Unsupported_PaymentApplication)(nil)            // 4
	_ MessageApplicationSub = (*waMultiDevice.MultiDevice)(nil)                 // 5
	_ MessageApplicationSub = (*Unsupported_Voip)(nil)                          // 6
	_ MessageApplicationSub = (*waArmadilloApplication.Armadillo)(nil)          // 7
)

func (*Unsupported_BusinessApplication) IsMessageApplicationSub() {}
func (*Unsupported_PaymentApplication) IsMessageApplicationSub()  {}
func (*Unsupported_Voip) IsMessageApplicationSub()                {}
