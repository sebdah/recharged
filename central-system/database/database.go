package database

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/dynamodb"
	"github.com/sebdah/recharged/central-system/settings"
)

var Db *dynamodb.DynamoDB

// Connect to AWS DynamoDB. Set region to "local" for DynamoDB Local
func GetDb() *dynamodb.DynamoDB {
	if Db == nil {
		conf := settings.GetSettings()
		credentials := aws.Creds(conf.AwsAccessKey, conf.AwsSecretKey, "")
		return dynamodb.New(&aws.Config{Region: conf.AwsRegion, Credentials: credentials})
	} else {
		return Db
	}
}
