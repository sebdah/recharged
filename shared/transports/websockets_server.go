package transports

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/sebdah/recharged/shared/processors"
)

type WsServer struct {
	Upgrader      websocket.Upgrader
	ConfProcessor processors.ConfProcessor
	ReqProcessor  processors.ReqProcessor
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

	go WsStreamReader(conn, this.GetReqProcessor(), this.GetConfProcessor())
}

// Get conf processor function
func (this *WsServer) GetConfProcessor() *processors.ConfProcessor {
	return &this.ConfProcessor
}

// Get req processor function
func (this *WsServer) GetReqProcessor() *processors.ReqProcessor {
	return &this.ReqProcessor
}

// Set conf processor function
func (this *WsServer) SetConfProcessor(proc processors.ConfProcessor) {
	this.ConfProcessor = proc
}

// Set req processor function
func (this *WsServer) SetReqProcessor(proc processors.ReqProcessor) {
	this.ReqProcessor = proc
}
