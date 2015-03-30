package database

import (
	"fmt"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/dynamodb"
	"github.com/sebdah/recharged/central-system/settings"
)

var Db *dynamodb.DynamoDB

// Connect to AWS DynamoDB. Set region to "local" for DynamoDB Local
func GetDb() *dynamodb.DynamoDB {
	if Db == nil {
		conf := settings.GetSettings()
		fmt.Printf("Connecting to DynamoDB in region '%s'\n", conf.AwsRegion)

		awsConf := &aws.Config{
			Region:      conf.AwsRegion,
			Credentials: aws.Creds(conf.AwsAccessKey, conf.AwsSecretKey, ""),
		}
		if conf.AwsDynamoDBEndpoint != "" {
			awsConf.Endpoint = conf.AwsDynamoDBEndpoint
		}

		return dynamodb.New(awsConf)
	} else {
		return Db
	}
}
