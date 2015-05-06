package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sebdah/recharged/central-system/database"
	"github.com/sebdah/recharged/central-system/models"
	"github.com/sebdah/recharged/central-system/settings"
	"github.com/sebdah/recharged/shared/websockets"
)

var WsServer *websockets.Server

func main() {
	log.Println("Starting OCPP central-system")

	// Set the environment
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}
	log.Printf("Using environment '%s'\n", env)

	// Get correct settings
	conf := settings.GetSettings()

	// Create databases if needed
	if env == "dev" {
		log.Println("Ensuring databases")
		database.CreateCollectionBootNotificationLog()

		log.Println("Ensuring indexes")
		models.EnsureIndexes(new(models.BootNotificationLog))
	}

	// Create the websockets server
	WsServer = websockets.NewServer()

	// Start the websockets message parser
	go messageParser()

	// Handlers
	http.HandleFunc("/ocpp-2.0j/ws", WsServer.Handler)

	// Start the HTTP server
	log.Printf("Starting webserver on port %d\n", conf.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), nil)
	if err != nil {
		panic(err)
	}
}

// Websockets message parser
func messageParser() {
	var recv_msg string

	for {
		recv_msg = <-WsServer.ReadMessage

		log.Printf("RECV: %s\n", recv_msg)
	}
}
