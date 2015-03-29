package types

import (
	"github.com/awslabs/aws-sdk-go/service/dynamodb"
	"github.com/sebdah/recharged/central-system/database"
)

var Db *dynamodb.DynamoDB = database.GetDb()

type IdTag struct {
	idTag IdToken
}

// Constructor
func NewIdTag(idTag IdToken) *IdTag {
	tag := new(IdTag)
	tag.idTag = idTag

	return tag
}

// Add to whitelist
func (idtag *IdTag) Whitelist() (bool, error) {
	var err error
	var result bool = false

	return result, err
}
