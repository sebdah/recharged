package types

import (
	"gopkg.in/mgo.v2/bson"
)

type IdTag struct {
	Id         bson.ObjectId `json:"-" bson:"_id,omitempty"`
	IdTag      string        `json:"idTag" type:"string" required:"true"`
	IdType     string        `json:"idType" type:"string" required:"true" default:"ISO14443"`
	ExpiryDate JSONTime      `json:"expiryDate" type:"Time" required:"false"`
	GroupIdTag string        `json:"groupIdTag" type:"string" required:"false"`
	Language   string        `json:"language" type:"string" required:"false" default:"en"`
	Active     bool          `json:"active" type:"bool" required:"true" default:"true"`
}

// Constructor
func NewIdTag() (idTag *IdTag) {
	idTag = new(IdTag)
	idTag.IdType = IdTypeISO14443
	idTag.Language = "en"
	idTag.Active = true

	return
}
