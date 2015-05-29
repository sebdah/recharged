package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	goLogging "github.com/op/go-logging"
	"github.com/sebdah/recharged/central-system/actions"
	"github.com/sebdah/recharged/central-system/config"
	"github.com/sebdah/recharged/central-system/heartbeat"
	"github.com/sebdah/recharged/central-system/logging"
	"github.com/sebdah/recharged/shared/messages"
	"github.com/sebdah/recharged/shared/rpc"
	"github.com/sebdah/recharged/shared/types"
	"github.com/sebdah/recharged/shared/websockets"
)

var (
	log      goLogging.Logger
	WsServer *websockets.Server
)

func main() {
	// Configure logging
	logging.Setup()

	// Welcome message
	log.Info("Starting re:charged central system service")
	log.Info("Environment: %s", config.Env)

	// Setup Websockets endpoint
	WsServer = websockets.NewServer()
	WsServer.SetPingNotificationChannel(heartbeat.HeartbeatChannel)

	// Fire up the websockets communicator
	go websocketCommunicator()

	// Configure handlers
	router := mux.NewRouter()
	router.HandleFunc("/ocpp-2.0j/ws/{chargePoint}", WsServer.Handler)
	http.Handle("/", router)

	// Start the HTTP server
	log.Info("Starting webserver on port %d", config.Config.GetInt("port"))
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Config.GetInt("port")), nil)
	if err != nil {
		panic(err)
	}
}

// Communicator for websockets, reading and sending messages
func websocketCommunicator() {
	var message string
	log.Info("Starting the websocket communicator")

	for {
		message = <-WsServer.ReadChannel
		log.Debug("Incoming message from %s: %s", WsServer.ChargePoint, message)
		messageType, err := rpc.ParseMessage(message)
		if err != nil {
			log.Notice("The incoming message does not match the RPC protocol")
			continue
		}

		switch {
		case messageType == 2: // Handle CALL
			var callError *rpc.CallError
			//var callResult rpc.CallResult

			log.Debug("Incoming RPC CALL: %s", message)

			// Instanciate the CALL
			call := rpc.NewCall()
			callError = call.Populate(message)
			if callError != nil {
				sendMessage(callError.String())
				continue
			}

			switch {
			case call.Action == "Authorize":
				// Populate the request
				authorizeReq, err := messages.NewAuthorizeReq(call.Payload)
				authorizeReq.SetChargePoint(WsServer.ChargePoint)
				if err != nil {
					callError = rpc.NewCallError(call.UniqueId, err)
					sendMessage(callError.String())
					continue
				}

				// Process the request
				authorizeConf, err := actions.Authorize(authorizeReq)
				if err != nil {
					callError = rpc.NewCallError(call.UniqueId, err)
					sendMessage(callError.String())
					continue
				}
				sendMessage(authorizeConf.String())
			case call.Action == "BootNotification":
				// Populate the request
				bootNotificationReq, err := messages.NewBootNotificationReq(call.Payload)
				bootNotificationReq.SetChargePoint(WsServer.ChargePoint)
				if err != nil {
					callError = rpc.NewCallError(call.UniqueId, err)
					sendMessage(callError.String())
					continue
				}

				// Process the request
				bootNotificationConf, err := actions.BootNotification(*bootNotificationReq)
				if err != nil {
					callError = rpc.NewCallError(call.UniqueId, err)
					sendMessage(callError.String())
					continue
				}
				sendMessage(bootNotificationConf.String())
			case call.Action == "DataTransfer":
				dataTransferConf := messages.NewDataTransferConf(types.DataTransferStatusRejected)
				sendMessage(dataTransferConf.String())
			}

		case messageType == 3: // Handle CALLRESULT
			log.Debug("Incoming RPC CALLRESULT: %s", message)
		case messageType == 4: // Handle CALLERROR
			log.Debug("Incoming RPC CALLERROR: %s", message)
		default:
			log.Error("RPC call not supported")
			continue
		}
	}
}

func sendMessage(msg string) {
	log.Debug("Sending message: %s", msg)
	WsServer.WriteChannel <- msg
}
