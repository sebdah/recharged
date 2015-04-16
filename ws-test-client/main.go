package main

import (
	"log"
	"net"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

var endpoint, _ = url.Parse("http://localhost:5000/ocpp/v2.0j/ws")
var wsHeaders = http.Header{
	"Origin": {"http://localhost:5000"},
	// your milage may differ
	"Sec-WebSocket-Extensions": {"permessage-deflate; client_max_window_bits, x-webkit-deflate-frame"},
}

func main() {
	// Create websockets connection
	_ = connectWebsockets()
	log.Printf("Connected to endpoint '%s' via websockets\n", endpoint)
}

// Create websockets connection
func connectWebsockets() (conn *websocket.Conn) {
	// Open new TCP socket
	rawConn, err := net.Dial("tcp", endpoint.Host)
	if err != nil {
		panic(err)
	}

	// Connect to websockets endpoint
	conn, _, err = websocket.NewClient(rawConn, endpoint, wsHeaders, 1024, 1024)
	if err != nil {
		panic(err)
	}

	return
}
