package database

import (
	"log"

	"github.com/awslabs/aws-sdk-go/service/dynamodb"
)

// Ensure that all tables exist
func EnsureTables(drop bool) {
	tables := Tables()
	listTablesInput := dynamodb.ListTablesInput{}

	listTablesOutput, err := Db.ListTables(&listTablesInput)
	if err != nil {
		log.Fatal(err)
	}

	// Drop the table if it exists
	if drop == true {
		for i := range listTablesOutput.TableNames {
			for j := range tables {
				if &tables[j].TableName == listTablesOutput.TableNames[i] {
					DropTable(tables[j])
				}
			}
		}
	}

	// Create the tables
	for i := range tables {
		CreateTable(tables[i])
	}
}

// Create DynamoDB table
func CreateTable(table Table) (output *dynamodb.CreateTableOutput, err error) {
	output, err = Db.CreateTable(&table.TableInput)
	if err != nil {
		log.Fatal(err)
	}
	return
}

// Drop a DynamoDB table
func DropTable(table Table) (bool, error) {
	deleteTableInput := dynamodb.DeleteTableInput{
		TableName: &table.TableName,
	}

	_, err := Db.DeleteTable(&deleteTableInput)
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	return true, err
}
