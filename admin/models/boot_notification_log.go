package models

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/sebdah/recharged/shared/types"
	"github.com/sebdah/recharged/admin/database"
)

type BootNotificationLog struct {
	Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Ts time.Time     `json:"ts"`

	*types.BootNotificationLog
}

// Constructor
func NewBootNotificationLog() (bootNotificationLog *BootNotificationLog) {
	bootNotificationLog = new(BootNotificationLog)
	bootNotificationLog.Ts = time.Now().UTC()

	return
}

// Get the collection, satisfies the Modeller interface
func (this *BootNotificationLog) Collection() *mgo.Collection {
	return database.GetCollectionBootNotificationLogs()
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
