package settings

import (
	"net/url"
	"os"
)

type Settings struct {
	AdminServiceUrl *url.URL `URL to the admin service`
	DatabaseName    string   `MongoDB database`
	MongoDBHosts    string   `MongoDB host:port combinations`
	Port            int64    `Service port number`
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
	settings.Port = 5000

	return
}
