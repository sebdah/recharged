package messages

import "github.com/sebdah/recharged/central-system/types"

type BootNotificationReq struct {
	ChargePointModel        string `json:"chargePointModel" type:"string" required:"true"`
	ChargePointVendor       string `json:"chargePointVendor" type:"string" required:"true"`
	ChargePointSerialNumber string `json:"chargePointSerialNumber" type:"string" required:"false"`
	Imsi                    string `json:"imsi" type:"string" required:"false"`
}

type BootNotificationConf struct {
	CurrentTime       types.JSONTime `json:"currentTime" type:"time.Time" required:"true"`
	HeartbeatInterval int64          `json:"heartbeatInterval" type:"int" required:"false"`
	Status            string         `json:"status" type:"string" required:"true"`
}
