package messages

type BootNotificationReq struct {
	ChargePointModel        string `json:"chargePointModel" type:"string" required:"true"`
	ChargePointVendor       string `json:"chargePointVendor" type:"string" required:"true"`
	ChargePointSerialNumber string `json:"chargePointSerialNumber" type:"string" required:"false"`
	Imsi                    string `json:"imsi" type:"string" required:"false"`
}

type BootNotificationConf struct {
	CurrentTime       typse.JSONTime `json:"currentTime" type:"time.Time" required:"true"`
	HeartbeatInterval int64          `json:"heartbeatInterval" type:"int" required:"false"`
	Status            bool           `json:"status" type:"bool" required:"true"`
}
