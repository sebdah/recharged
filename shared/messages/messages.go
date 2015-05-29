package messages

import "github.com/op/go-logging"

var log logging.Logger

// Base structure for all requests
type BaseReq struct {
	ChargePoint string
}

// Set ChargePoint
func (this *BaseReq) SetChargePoint(cp string) {
	this.ChargePoint = cp
}