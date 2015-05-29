package types

type BootNotificationLog struct {
	ChargePoint ChargePoint `json:"chargePoint"`
}

// Constructor
func NewBootNotificationLog() (bootNotificationLog *BootNotificationLog) {
	bootNotificationLog = new(BootNotificationLog)
	return
}
