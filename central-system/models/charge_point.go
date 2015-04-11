package models

import (
	"github.com/sebdah/recharged/central-system/database"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

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

// Get the collection, satisfies the Modeller interface
func (this *ChargePoint) Collection() *mgo.Collection {
	return database.GetCollectionChargePoints()
}

// Indexes, satisfies the Modeller interface
func (this *ChargePoint) Indexes() (indexes []*mgo.Index) {
	imsiIndex := mgo.Index{
		Key:    []string{"imsi"},
		Unique: true,
	}

	indexes = append(indexes, &index)

	return
}

// Get the ID, satisfies the Modeller interface
func (this *ChargePoint) GetId() bson.ObjectId {
	return this.Id
}

// Set the ID, satisfies the Modeller interface
func (this *ChargePoint) SetId(id *bson.ObjectId) {
	this.Id = *id
}
