package websockets

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn            *websocket.Conn
	Endpoint        *url.URL
	Headers         http.Header
	ReadBufferSize  int
	WriteBufferSize int
	WriteChannel    chan string
	ReadChannel     chan string
	PingChannel     chan string
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
	client.ReadChannel = make(chan string)
	client.WriteChannel = make(chan string)
	client.PingChannel = make(chan string)

	client.connect()

	return
}

// Connect to a websockets server
func (this *Client) connect() {
	rawConn, err := net.Dial("tcp", this.Endpoint.Host)
	if err != nil {
		panic(err)
	}

	this.Conn, _, err = websocket.NewClient(
		rawConn,
		this.Endpoint,
		this.Headers,
		this.ReadBufferSize,
		this.WriteBufferSize)
	if err != nil {
		panic(err)
	}

	// Set some limits
	this.Conn.SetReadLimit(maxMessageSize)
	this.Conn.SetReadDeadline(time.Now().Add(pongWait))

	// Register pong handler
	this.Conn.SetPongHandler(func(message string) error {
		log.Debug("Received pong message")
		return nil
	})

	log.Info("Connected to endpoint '%s' via websockets\n", this.Endpoint)

	// Instanciate a new communicator
	communicator := NewCommunicator(this.Conn)
	log.Debug("Starting websockets communication channel")
	go communicator.Pinger(this.PingChannel)
	go communicator.Reader(this.ReadChannel)
	go communicator.Writer(this.WriteChannel)
}

// Send ping message
func (this *Client) SendPing(message string) {
	this.PingChannel <- message
}
