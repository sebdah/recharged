package settings

import "os"

type Settings struct {
	MongoDBHosts string
	DatabaseName string
}

func GetSettings() Settings {
	devSettings := Settings{
		MongoDBHosts: "localhost:27017",
		DatabaseName: "rechargedDev",
	}

	environment := os.Getenv("ENV")

	switch {
	case environment == "dev":
		return devSettings
	default:
		return devSettings
	}
}
