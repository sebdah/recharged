package main

import (
	"fmt"
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
	fmt.Printf("Using environment '%s'\n", env)

	// Set default port
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	// Create databases if needed
	if env == "dev" {
		fmt.Println("Ensuring databases")
		database.CreateCollectionIdTags()
	}

	idTag := models.NewIdTag("id")
	idTag.IdTag = "id"
	models.Save(idTag)
	idTag.IdTag = "newid"
	models.Save(idTag)

	fmt.Printf("Starting webserver on port %s\n", port)
	http.ListenAndServe(":"+port, routers.Router())
}
