package transports

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/gorilla/websocket"
	"github.com/sebdah/recharged/central-system/messages"
	"github.com/sebdah/recharged/central-system/rpc"
)

var messageTypeRegExp = regexp.MustCompile(`^\[(?P<messageId>\d+),(.*)\]$`)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// Handler registering connections
func WebsocketTransport(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading websocket connection: %s\n", err)
		return
	}

	go websocketStreamReader(conn)
}

// Message loop
func websocketStreamReader(conn *websocket.Conn) {
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
			call, callError := callHandler(string(msg))

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

// Handle CALL requests
func callHandler(msg string) (callResult *rpc.CallResult, callError *rpc.CallError) {
	var err error

	// Populate and validate the fields
	call := rpc.NewCall()
	callError = call.Populate(msg)
	if callError != nil {
		return
	}

	// Find the requested message matching the action
	if call.Action == "Authorize" {
		req := messages.NewAuthorizeReq()
		err = req.ParseJson(call.Payload)
		if err != nil {
			log.Println("Error parsing JSON: %s", err.Error())
			callError = rpc.NewCallError(call.UniqueId, rpc.NewFormationViolation())
			return
		}
		conf, err := req.Process()
		if err != nil {
			log.Println("Error processing request: %s", err.Error())
			genericError := rpc.NewGenericError()
			genericError.SetDetails(fmt.Sprintf(`{"message": "%s"}`, err.Error()))
			callError = rpc.NewCallError(call.UniqueId, genericError)
			return
		}

		callResult = rpc.NewCallResult(call.UniqueId)
		callResult.SetPayload(conf.String())
		return

	} else if call.Action == "BootNotification" {
		req := messages.NewBootNotificationReq()
		err = req.ParseJson(call.Payload)
		if err != nil {
			log.Println("Error parsing JSON: %s", err.Error())
			callError = rpc.NewCallError(call.UniqueId, rpc.NewFormationViolation())
			return
		}
		conf, err := req.Process()
		if err != nil {
			log.Println("Error processing request: %s", err.Error())
			genericError := rpc.NewGenericError()
			genericError.SetDetails(fmt.Sprintf(`{"message": "%s"}`, err.Error()))
			callError = rpc.NewCallError(call.UniqueId, genericError)
			return
		}

		callResult = rpc.NewCallResult(call.UniqueId)
		callResult.SetPayload(conf.String())
		return

	} else if call.Action == "DataTransfer" {
		req := messages.NewDataTransferReq()
		err = req.ParseJson(call.Payload)
		if err != nil {
			log.Println("Error parsing JSON: %s", err.Error())
			callError = rpc.NewCallError(call.UniqueId, rpc.NewFormationViolation())
			return
		}
		conf, err := req.Process()
		if err != nil {
			log.Println("Error processing request: %s", err.Error())
			genericError := rpc.NewGenericError()
			genericError.SetDetails(fmt.Sprintf(`{"message": "%s"}`, err.Error()))
			callError = rpc.NewCallError(call.UniqueId, genericError)
			return
		}

		callResult = rpc.NewCallResult(call.UniqueId)
		callResult.SetPayload(conf.String())
		return

	} else {
		callError = rpc.NewCallError(call.UniqueId, rpc.NewNotImplementedError())
		return
	}

	return
}
