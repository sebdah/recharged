package database

import (
	"fmt"
	"log"

	"github.com/awslabs/aws-sdk-go/service/dynamodb"
)

type Table struct {
	TableName          string
	ReadCapacityUnits  int64
	WriteCapacityUnits int64
	HashKeyName        string
	HashKeyType        string
	RangeKeyName       string
	RangeKeyType       string
}

// Create table
func (table *Table) CreateTable() (output *dynamodb.CreateTableOutput, err error) {
	fmt.Printf("Creating table '%s'.. ", table.TableName)
	input := table.GetCreateTableInput()
	output, err = Db.CreateTable(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("OK")
	return
}

// Delete table
func (table *Table) DeleteTable() (output *dynamodb.DeleteTableOutput, err error) {
	fmt.Printf("Deleting table '%s'.. ", table.TableName)
	input := table.GetDeleteTableInput()
	output, err = Db.DeleteTable(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("OK")
	return
}

// Ensure table
func (table *Table) EnsureTable() {
	tables, err := Db.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		log.Fatal(err)
	}

	// Skip table creation if the table already exists
	for i := range tables.TableNames {
		if tables.TableNames[i] == &table.TableName {
			return
		}
	}

	// Create the table
	table.CreateTable()
	return
}

// Define the table for creation
func (table *Table) GetCreateTableInput() *dynamodb.CreateTableInput {
	var (
		hashKeyType  string = "HASH"
		rangeKeyType string = "RANGE"
	)

	// Create provisioning settings
	provisionedThroughput := dynamodb.ProvisionedThroughput{
		ReadCapacityUnits:  &table.ReadCapacityUnits,
		WriteCapacityUnits: &table.WriteCapacityUnits,
	}

	// Generate attribute definitions
	var attributeDefinitions []*dynamodb.AttributeDefinition
	hashAttributeDefinition := &dynamodb.AttributeDefinition{
		AttributeName: &table.HashKeyName,
		AttributeType: &table.HashKeyType,
	}
	attributeDefinitions = append(attributeDefinitions, hashAttributeDefinition)

	if table.RangeKeyName != "" {
		rangeAttributeDefinition := &dynamodb.AttributeDefinition{
			AttributeName: &table.RangeKeyName,
			AttributeType: &table.RangeKeyType,
		}
		attributeDefinitions = append(attributeDefinitions, rangeAttributeDefinition)
	}

	// Generate key schema
	var keySchema []*dynamodb.KeySchemaElement
	hashKey := &dynamodb.KeySchemaElement{
		AttributeName: &table.HashKeyName,
		KeyType:       &hashKeyType,
	}
	keySchema = append(keySchema, hashKey)
	if table.RangeKeyName != "" {
		rangeKey := &dynamodb.KeySchemaElement{
			AttributeName: &table.RangeKeyName,
			KeyType:       &rangeKeyType,
		}
		keySchema = append(keySchema, rangeKey)
	}

	return &dynamodb.CreateTableInput{
		AttributeDefinitions:  attributeDefinitions,
		KeySchema:             keySchema,
		ProvisionedThroughput: &provisionedThroughput,
		TableName:             &table.TableName,
	}
}

// Define the table for deletion
func (table *Table) GetDeleteTableInput() (input *dynamodb.DeleteTableInput) {
	input = &dynamodb.DeleteTableInput{
		TableName: &table.TableName,
	}

	return
}

// Recreate the table (or just create it if it doesn't exist)
func (table *Table) RecreateTable() {
	tables, err := Db.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		log.Fatal(err)
	}

	// Delete table if the table already exists
	for i := range tables.TableNames {
		if tables.TableNames[i] == &table.TableName {
			table.DeleteTable()
		}
	}

	// Create the table
	table.CreateTable()
	return
}
