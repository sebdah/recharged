package messages

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/sebdah/recharged/shared/rpc"
	"github.com/sebdah/recharged/shared/types"
)

type BootNotificationReq struct {
	messageType             string `json:"-" type:"string"`
	ChargePointModel        string `json:"chargePointModel" type:"string" required:"true"`
	ChargePointVendor       string `json:"chargePointVendor" type:"string" required:"true"`
	ChargePointSerialNumber string `json:"chargePointSerialNumber" type:"string" required:"false"`
	Imsi                    string `json:"imsi" type:"string" required:"false"`

	*BaseReq
}

type BootNotificationConf struct {
	CurrentTime       types.JSONTime `json:"currentTime" type:"time.Time" required:"true"`
	HeartbeatInterval int64          `json:"heartbeatInterval" type:"int" required:"false"`
	Status            string         `json:"status" type:"string" required:"true"`
}

func NewBootNotificationReq(payload string) (req *BootNotificationReq, rpcError rpc.Errorer) {
	req = new(BootNotificationReq)
	req.messageType = "BootNotification"

	decoder := json.NewDecoder(strings.NewReader(payload))
	err := decoder.Decode(&req)
	if err != nil {
		log.Notice("Unable to parse payload: %s", err.Error())
		rpcError = rpc.NewFormationViolation()
		return
	}

	return
}

func NewBootNotificationConf(status string) (conf *BootNotificationConf) {
	conf = new(BootNotificationConf)
	conf.CurrentTime.Time = time.Now()
	conf.HeartbeatInterval = 60
	conf.Status = status
	return
}

// Get the message type
func (this *BootNotificationReq) GetMessageType() string {
	return this.messageType
}

// Populate the object with JSON data
func (this *BootNotificationReq) ParseJson(data string) (err error) {
	decoder := json.NewDecoder(strings.NewReader(data))
	err = decoder.Decode(&this)
	if err != nil {
		log.Error("Unable to parse payload: %s", err.Error())
		return
	}

	return
}

// String representation
func (this *BootNotificationReq) String() (str string) {
	js, _ := json.Marshal(this)
	str = string(js)
	return
}

// String representation
func (this *BootNotificationConf) String() (str string) {
	js, _ := json.Marshal(this)
	str = string(js)
	return
}
