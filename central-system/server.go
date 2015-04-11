package main

import (
	"log"
	"net/http"
	"os"

	"github.com/sebdah/recharged/central-system/database"
	"github.com/sebdah/recharged/central-system/models"
	"github.com/sebdah/recharged/central-system/routers"
)

func main() {
	// Set the environment
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}
	log.Printf("Using environment '%s'\n", env)

	// Set default port
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	// Create databases if needed
	if env == "dev" {
		log.Println("Ensuring databases")
		database.CreateCollectionIdTags()
		database.CreateCollectionChargePoints()

		log.Println("Ensuring indexes")
		models.EnsureIndexes(new(models.IdTag))
		models.EnsureIndexes(new(models.ChargePoint))
	}

	log.Printf("Starting webserver on port %s\n", port)
	http.ListenAndServe(":"+port, routers.Router())
}
