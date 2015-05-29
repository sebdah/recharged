package types

type ChargePoint struct {
	Id           string `json:"id" type:"string" required:"true"`
	Model        string `json:"model" type:"string" required:"true"`
	Vendor       string `json:"vendor" type:"string" required:"true"`
	SerialNumber string `json:"serialNumber" type:"string" required:"false"`
	Imsi         string `json:"imsi" type:"string" required:"false"`
}

// Constructor
func NewChargePoint() (chargePoint *ChargePoint) {
	chargePoint = new(ChargePoint)
	return
}
