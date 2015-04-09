package models

import (
	"github.com/sebdah/recharged/central-system/database"
	"github.com/sebdah/recharged/central-system/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type IdTag struct {
	Id         bson.ObjectId  `json:"-" bson:"_id,omitempty"`
	IdTag      string         `json:"idTag" type:"string" required:"true"`
	IdType     string         `json:"idType" type:"string" required:"true" default:"ISO14443"`
	ExpiryDate types.JSONTime `json:"expiryDate" type:"Time" required:"false"`
	GroupIdTag string         `json:"groupIdTag" type:"string" required:"false"`
	Language   string         `json:"language" type:"string" required:"false" default:"en"`
	Active     bool           `json:"-" type:"bool" required:"true" default:"false"`
}

// Constructor
func NewIdTag() (idTag *IdTag) {
	idTag = new(IdTag)
	idTag.IdType = types.IdTypeISO14443
	idTag.Language = "en"
	idTag.Active = false

	return
}

// Get the collection, satisfies the Modeller interface
func (this *IdTag) Collection() *mgo.Collection {
	return database.GetCollectionIdTags()
}

// Indexes, satisfies the Modeller interface
func (this *IdTag) Indexes() (indexes []*mgo.Index) {
	idTagIndex := mgo.Index{
		Key:    []string{"idtag"},
		Unique: true,
	}
	indexes = append(indexes, &idTagIndex)

	return
}

// Get the ID, satisfies the Modeller interface
func (this *IdTag) GetId() bson.ObjectId {
	return this.Id
}

// Set the ID, satisfies the Modeller interface
func (this *IdTag) SetId(id *bson.ObjectId) {
	this.Id = *id
}
