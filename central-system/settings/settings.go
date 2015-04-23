package settings

import (
	"net/url"
	"os"
)

type Settings struct {
	AdminServiceUrl *url.URL
	DatabaseName    string
	MongoDBHosts    string
}

func GetSettings() Settings {
	environment := os.Getenv("ENV")

	switch {
	case environment == "dev":
		return GetDevSettings()
	default:
		return GetDevSettings()
	}
}

func GetDevSettings() (settings Settings) {
	adminServiceUrl, _ := url.Parse("http://localhost:6000")

	settings.AdminServiceUrl = adminServiceUrl
	settings.DatabaseName = "rechargedDevCs"
	settings.MongoDBHosts = "localhost:27017"

	return
}
