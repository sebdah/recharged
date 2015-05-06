package websockets

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Server struct {
	Upgrader        websocket.Upgrader
	ReadBufferSize  int
	WriteBufferSize int
	WriteMessage    chan string
	ReadMessage     chan string
}

func NewServer() (server *Server) {
	server = new(Server)
	server.ReadBufferSize = 1024
	server.WriteBufferSize = 1024

	server.Upgrader = websocket.Upgrader{
		ReadBufferSize:  server.ReadBufferSize,
		WriteBufferSize: server.WriteBufferSize,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	// Create channels
	server.ReadMessage = make(chan string)
	server.WriteMessage = make(chan string)

	return
}

// Handler registering connections
func (this *Server) Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := this.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading websocket connection: %s\n", err)
		return
	}

	// Instanciate a new communicator
	communicator := NewCommunicator(conn)
	communicator.Name = "Server"
	log.Println("Starting websockets communication channel")
	go communicator.Reader(this.ReadMessage)
	go communicator.Writer(this.WriteMessage)
}
