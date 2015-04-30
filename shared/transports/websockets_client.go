package transports

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

type WsClient struct {
	Endpoint *url.URL
	Headers  http.Header
	Conn     *websocket.Conn
}

// Constructor
func NewWsClient(endpoint *url.URL) (wsClient *WsClient) {
	wsClient = new(WsClient)
	wsClient.Endpoint = endpoint
	wsClient.Headers = http.Header{
		"Origin":                   {fmt.Sprintf("%s://%s", endpoint.Scheme, endpoint.Host)},
		"Sec-WebSocket-Extensions": {"permessage-deflate; client_max_window_bits, x-webkit-deflate-frame"},
	}

	return
}

// Connect to a websockets server
func (this *WsClient) Connect() {

	rawConn, err := net.Dial("tcp", this.Endpoint.Host)
	if err != nil {
		panic(err)
	}

	conn, _, err := websocket.NewClient(rawConn, this.Endpoint, this.Headers, 1024, 1024)
	if err != nil {
		panic(err)
	}

	log.Printf("Connected to endpoint '%s' via websockets\n", this.Endpoint)
	this.Conn = conn

	go WsStreamReader(conn, reqProcessor, confProcessor)
}

// Send message to the stream
func (this *WsClient) SendMessage(msg string) {
	this.Conn.WriteMessage(websocket.TextMessage, msg)
}
