package models

import (
	"github.com/sebdah/recharged/central-system/database"
	"github.com/sebdah/recharged/central-system/types"
)

var table Table = database.TableIdTags()

type IdTag struct {
	Id         string `type:"string" required:"true"`
	IdType     string `type:"string" required:"true" default:"ISO14443"`
	ExpiryDate string `type:"string" required:"false"`
	GroupIdTag string `type:"string" required:"false"`
	Language   string `type:"string" required:"false" default:"en"`
	Active     bool   `type:"bool" required:"true" default:"false"`
}

// Constructor
func NewIdTag(id string) (idTag *IdTag) {
	idTag.Id = id
	idTag.idType = types.IdTypeISO14443
	idTag.Language = "en"
	idTag.Active = false
	return
}
