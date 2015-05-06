package websockets

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

type Client struct {
	Endpoint        *url.URL
	Headers         http.Header
	ReadBufferSize  int
	WriteBufferSize int
	WriteMessage    chan string
	ReadMessage     chan string
}

// Constructor
func NewClient(endpoint *url.URL) (client *Client) {
	client = new(Client)
	client.Endpoint = endpoint
	client.ReadBufferSize = 1024
	client.WriteBufferSize = 1024
	client.Headers = http.Header{
		"Origin":                   {fmt.Sprintf("%s://%s", endpoint.Scheme, endpoint.Host)},
		"Sec-WebSocket-Extensions": {"permessage-deflate; client_max_window_bits, x-webkit-deflate-frame"},
	}

	// Create channels
	client.ReadMessage = make(chan string)
	client.WriteMessage = make(chan string)

	client.connect()

	return
}

// Connect to a websockets server
func (this *Client) connect() {
	rawConn, err := net.Dial("tcp", this.Endpoint.Host)
	if err != nil {
		panic(err)
	}

	conn, _, err := websocket.NewClient(
		rawConn,
		this.Endpoint,
		this.Headers,
		this.ReadBufferSize,
		this.WriteBufferSize)
	if err != nil {
		panic(err)
	}

	log.Printf("Connected to endpoint '%s' via websockets\n", this.Endpoint)

	// Instanciate a new communicator
	communicator := NewCommunicator(conn)
	communicator.Name = "Client"
	log.Println("Starting websockets communication channel")
	go communicator.Reader(this.ReadMessage)
	go communicator.Writer(this.WriteMessage)

}
