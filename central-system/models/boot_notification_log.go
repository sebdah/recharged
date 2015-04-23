package models

import (
	"time"

	"github.com/sebdah/recharged/central-system/database"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type BootNotificationLog struct {
	Id           bson.ObjectId `bson:"_id,omitempty"`
	Vendor       string
	Model        string
	SerialNumber string
	Imsi         string
	Ts           time.Time
}

// Constructor
func NewBootNotificationLog() (bootNotificationLog *BootNotificationLog) {
	bootNotificationLog = new(BootNotificationLog)
	bootNotificationLog.Ts = time.Now().UTC()

	return
}

// Get the collection, satisfies the Modeller interface
func (this *BootNotificationLog) Collection() *mgo.Collection {
	return database.GetCollectionBootNotificationLog()
}

// Indexes, satisfies the Modeller interface
func (this *BootNotificationLog) Indexes() (indexes []*mgo.Index) {
	index := mgo.Index{}
	indexes = append(indexes, &index)

	return
}

// Get the ID, satisfies the Modeller interface
func (this *BootNotificationLog) GetId() bson.ObjectId {
	return this.Id
}

// Set the ID, satisfies the Modeller interface
func (this *BootNotificationLog) SetId(id *bson.ObjectId) {
	this.Id = *id
}
