package main

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/sebdah/recharged/ws/shared/websockets"
)

func main() {
	log.Println("Starting websockets client")
	endpoint, err := url.Parse("http://localhost:5000/ws")
	if err != nil {
		panic(err)
	}

	wsClient := websockets.NewClient(endpoint)

	// Send pings
	go func() {
		for {
			wsClient.WriteMessage <- "ping"
			time.Sleep(1000 * time.Millisecond)
		}
	}()

	// Read incoming messages
	go func() {
		var recv_msg string

		for {
			recv_msg = <-wsClient.ReadMessage

			// Craft the response
			if string(recv_msg) == "ping" {
				wsClient.WriteMessage <- "pong"
			} else if string(recv_msg) == "pong" {
				wsClient.WriteMessage <- ""
			} else {
				wsClient.WriteMessage <- "unhandled message"
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
