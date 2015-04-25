package transports

import (
	"regexp"

	"github.com/gorilla/websocket"
	"github.com/sebdah/recharged/shared/handlers"
	"github.com/sebdah/recharged/shared/processors"
	"github.com/sebdah/recharged/shared/rpc"
)

// Message loop
func WsStreamReader(conn *websocket.Conn, reqProcessor *processors.ReqProcessor) {
	messageTypeRegExp := regexp.MustCompile(`^\[(?P<messageId>\d+),(.*)\]$`)

	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		match := messageTypeRegExp.FindStringSubmatch(string(msg))

		// Handle malformatted calls
		if len(match) == 0 {
			callError := rpc.NewCallError("0", rpc.NewGenericError())

			err = conn.WriteMessage(messageType, []byte(callError.String()))
			if err != nil {
				return
			}
		}

		// Get the messageId
		result := make(map[string]string)
		for i, name := range messageTypeRegExp.SubexpNames() {
			result[name] = match[i]
		}

		// Handle CALL requests
		if result["messageId"] == "2" {
			call, callError := handlers.CallHandler(string(msg), *reqProcessor)

			// Send the response
			var msg string
			if callError != nil {
				msg = callError.String()
			} else {
				msg = call.String()
			}
			err = conn.WriteMessage(messageType, []byte(msg))
			if err != nil {
				return
			}
		}
	}
}
