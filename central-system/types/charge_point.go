package types

import "gopkg.in/mgo.v2/bson"

type ChargePoint struct {
	Id           bson.ObjectId `json:"-" bson:"_id,omitempty"`
	Model        string        `json:"model" type:"string" required:"true"`
	Vendor       string        `json:"vendor" type:"string" required:"true"`
	SerialNumber string        `json:"serialNumber" type:"string" required:"false"`
	Imsi         string        `json:"imsi" type:"string" required:"false"`
}

// Constructor
func NewChargePoint() (chargePoint *ChargePoint) {
	chargePoint = new(ChargePoint)
	return
}
