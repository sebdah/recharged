package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sebdah/recharged/admin/database"
	"github.com/sebdah/recharged/admin/models"
	"github.com/sebdah/recharged/admin/routers"
	"github.com/sebdah/recharged/admin/settings"
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
		database.CreateCollectionIdTags()
		database.CreateCollectionChargePoints()

		log.Println("Ensuring indexes")
		models.EnsureIndexes(new(models.IdTag))
		models.EnsureIndexes(new(models.ChargePoint))
	}

	log.Printf("Starting webserver on port %d\n", conf.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), routers.Router())
}
