package transports

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/sebdah/recharged/shared/processor"
)

type WsServer struct {
	Upgrader  websocket.Upgrader
	Processor processor.Processor
}

func NewWsServer() (server *WsServer) {
	server = new(WsServer)
	server.Upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	return
}

// Handler registering connections
func (this *WsServer) Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := this.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading websocket connection: %s\n", err)
		return
	}

	go WsStreamReader(conn, this.GetProcessor())
}

// Get processor function
func (this *WsServer) GetProcessor() *processor.Processor {
	return &this.Processor
}

// Set processor function
func (this *WsServer) SetProcessor(proc processor.Processor) {
	this.Processor = proc
}
