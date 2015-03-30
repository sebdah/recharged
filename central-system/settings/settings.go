package settings

import "os"

type Settings struct {
	AwsRegion           string
	AwsAccessKey        string
	AwsSecretKey        string
	AwsDynamoDBEndpoint string
}

func GetSettings() Settings {
	devSettings := Settings{
		AwsRegion:           "local",
		AwsAccessKey:        "foo",
		AwsSecretKey:        "bar",
		AwsDynamoDBEndpoint: "http://localhost:8000",
	}

	environment := os.Getenv("ENV")

	switch {
	case environment == "dev":
		return devSettings
	default:
		return devSettings
	}
}
