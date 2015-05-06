package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/sebdah/recharged/shared/websockets"
)

var WsClient *websockets.Client

func main() {
	log.Println("Starting OCPP charge-point")

	// Create websockets client
	centralSystemUrl, _ := url.Parse("http://localhost:5000/ocpp-2.0j/ws")
	WsClient = websockets.NewClient(centralSystemUrl)

	// Start the message parser
	go messageParser()

	// Start the dummy sender
	go func() {
		WsClient.WriteMessage <- "ping"
	}()

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}

// Websockets message parser
func messageParser() {
	var recv_msg string

	for {
		recv_msg = <-WsClient.ReadMessage

		log.Printf("RECV: %s\n", recv_msg)
	}
}
