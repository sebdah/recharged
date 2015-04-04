package models

import (
	"fmt"

	"github.com/sebdah/recharged/central-system/database"
	"github.com/sebdah/recharged/central-system/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var collection *mgo.Collection

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

	collection = database.GetCollectionIdTags()

	return
}

// Save the document
func (idTag *IdTag) Save() *IdTag {
	fmt.Printf("Id: %s\n", idTag.Id)
	if idTag.Id == "" {
		idTag.Id = bson.NewObjectId()
		collection.Insert(&idTag)
	} else {
		collection.Update(bson.M{"_id": idTag.Id}, &idTag)
	}

	return idTag
}
