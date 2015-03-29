package settings

import "os"

type Settings struct {
	AwsRegion    string
	AwsAccessKey string
	AwsSecretKey string
}

func GetSettings() Settings {
	devSettings := Settings{
		AwsRegion:    "local",
		AwsAccessKey: "foo",
		AwsSecretKey: "bar"}

	environment := os.Getenv("ENV")

	switch {
	case environment == "dev":
		return devSettings
	default:
		return devSettings
	}
}
