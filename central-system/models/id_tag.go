package models

import (
	"github.com/sebdah/recharged/central-system/database"
	"github.com/sebdah/recharged/central-system/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type IdTag struct {
	Id         bson.ObjectId `bson:"_id,omitempty"`
	IdTag      string        `type:"string" required:"true"`
	IdType     string        `type:"string" required:"true" default:"ISO14443"`
	ExpiryDate string        `type:"string" required:"false"`
	GroupIdTag string        `type:"string" required:"false"`
	Language   string        `type:"string" required:"false" default:"en"`
	Active     bool          `type:"bool" required:"true" default:"false"`
}

// Constructor
func NewIdTag(id string) (idTag *IdTag) {
	idTag = new(IdTag)
	idTag.IdTag = id
	idTag.IdType = types.IdTypeISO14443
	idTag.Language = "en"
	idTag.Active = false

	return
}

// Get the collection, satisfies the Modeller interface
func (this *IdTag) Collection() *mgo.Collection {
	return database.GetCollectionIdTags()
}

// Get the ID
func (this *IdTag) GetId() bson.ObjectId {
	return this.Id
}

// Set the ID
func (this *IdTag) SetId(id *bson.ObjectId) {
	this.Id = *id
}
