package main

import (
	"log"
	"net/http"

	"github.com/sebdah/recharged/ws/shared/websockets"
)

func main() {
	wsServer := websockets.NewServer()

	log.Println("Starting webserver on port 5000")
	http.HandleFunc("/ws", wsServer.Handler)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		panic(err)
	}
}
