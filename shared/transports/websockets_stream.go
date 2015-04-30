package transports

import (
	"github.com/sebdah/recharged/shared/rpc"

	"github.com/gorilla/websocket"
)

func WsStream(conn *websocket.Conn, c_sendCall chan rpc.Call) {
	var call rpc.Call

	for {
		// Read message from the stream
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

	}

}
