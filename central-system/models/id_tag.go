package models

import (
	"github.com/awslabs/aws-sdk-go/service/dynamodb"
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

// Save the item
func (idtag *IdTag) Save() bool {
	var item map[string]*dynamodb.AttributeValue
	var input dynamodb.PutItemInput

	idAttribute := new(dynamodb.AttributeValue)
	idAttribute.S = *idtag.Id
	idTypeAttribute := new(dynamodb.AttributeValue)
	idTypeAttribute.S = *idtag.IdType
	expiryDateAttribute := new(dynamodb.AttributeValue)
	expiryDateAttribute.S = *idtag.ExpiryDate
	groupTagIdAttribute := new(dynamodb.AttributeValue)
	groupTagIdAttribute.S = *idtag.GroupIdTag
	languageAttribute := new(dynamodb.AttributeValue)
	languageAttribute.S = *idtag.Language
	activeAttribute := new(dynamodb.AttributeValue)
	activeAttribute.B = *idtag.Active

	input.Item = *item

	_, err := table.PutItem(*input)
	if err != nil {
		return false
	}
	return true
}
