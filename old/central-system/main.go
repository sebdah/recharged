package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sebdah/recharged/central-system/database"
	"github.com/sebdah/recharged/central-system/models"
	"github.com/sebdah/recharged/central-system/routers"
	"github.com/sebdah/recharged/central-system/settings"
)

func main() {
	// Set the environment
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}
	log.Printf("Using environment '%s'\n", env)

	conf := settings.GetSettings()

	// Create databases if needed
	if env == "dev" {
		log.Println("Ensuring databases")
		database.CreateCollectionBootNotificationLog()

		log.Println("Ensuring indexes")
		models.EnsureIndexes(new(models.BootNotificationLog))
	}

	log.Printf("Starting webserver on port %d\n", conf.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), routers.Router())
}
