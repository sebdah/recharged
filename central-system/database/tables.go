package database

import "github.com/awslabs/aws-sdk-go/service/dynamodb"

type Table struct {
	TableName  string
	TableInput dynamodb.CreateTableInput
}

// Tables defined in the slice below will be used for
// table creation and removal
func Tables() (tables []Table) {
	tables = append(tables, Table{
		TableName:  "idtag_whitelist",
		TableInput: TableInputIdTagWhitelist(),
	})

	return
}
