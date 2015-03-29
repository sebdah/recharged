package database

import (
	"github.com/awslabs/aws-sdk-go/service/dynamodb"
)

func TableInputIdTagWhitelist() dynamodb.CreateTableInput {
	var (
		tableName     string = "idtag_whitelist"
		idTagName     string = "idtag"
		idTagType     string = "S"
		idTagKeyType  string = "HASH"
		readCapacity  int64  = 1
		writeCapacity int64  = 1
	)

	// Set provisioning
	provisionedThroughput := dynamodb.ProvisionedThroughput{
		ReadCapacityUnits:  &readCapacity,
		WriteCapacityUnits: &writeCapacity,
	}

	// Create list of attribute definitions
	var attributeDefinitions []*dynamodb.AttributeDefinition
	attributeDefinitions = append(attributeDefinitions, &dynamodb.AttributeDefinition{
		AttributeName: &idTagName,
		AttributeType: &idTagType,
	})

	// Create list of keys
	var keySchema []*dynamodb.KeySchemaElement
	keySchema = append(keySchema, &dynamodb.KeySchemaElement{
		AttributeName: &idTagName,
		KeyType:       &idTagKeyType,
	})

	return dynamodb.CreateTableInput{
		AttributeDefinitions:  attributeDefinitions,
		KeySchema:             keySchema,
		ProvisionedThroughput: &provisionedThroughput,
		TableName:             &tableName,
	}
}
