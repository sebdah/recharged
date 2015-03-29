package database

import "github.com/awslabs/aws-sdk-go/service/dynamodb"

type Table struct {
	TableName  string
	TableInput dynamodb.CreateTableInput
}

func Tables() (tables []Table) {
	tables = append(tables, Table{
		TableName:  "idtag_whitelist",
		TableInput: TableInputIdTagWhitelist(),
	})

	return
}
