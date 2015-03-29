package types

import (
	"log"

	"github.com/awslabs/aws-sdk-go/service/dynamodb"
	"github.com/sebdah/recharged/central-system/database"
)

var Db *dynamodb.DynamoDB = database.GetDb()

type IdTag struct {
	idTag IdToken
}

var TableName string = "idtag_whitelist"

// Constructor
func NewIdTag(idTag IdToken) *IdTag {
	tag := new(IdTag)
	tag.idTag = idTag

	return tag
}

// Add to whitelist
func (idtag *IdTag) AddToWhiteList() (bool, error) {
	var item map[string]*dynamodb.AttributeValue
	item["idtag"] = &dynamodb.AttributeValue{S: idtag.idTag.ToString()}

	putItemInput := dynamodb.PutItemInput{
		Item:      &item,
		TableName: &TableName,
	}

	_, err := Db.PutItem(&putItemInput)
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	return true, err
}

// Remove from whitelist
func (idtag *IdTag) RemoveFromWhitelist() (bool, error) {
	var item map[string]*dynamodb.AttributeValue
	item["idtag"] = &dynamodb.AttributeValue{S: idtag.idTag.ToString()}

	deleteItemInput := dynamodb.DeleteItemInput{
		Key:       &item,
		TableName: &TableName,
	}

	_, err := Db.DeleteItem(&deleteItemInput)
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	return true, err
}
