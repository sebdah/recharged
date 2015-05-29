package main

import (
	"flag"
	"fmt"
	"net/url"
	"time"

	goLogging "github.com/op/go-logging"
	"github.com/sebdah/recharged/charge-point/config"
	"github.com/sebdah/recharged/charge-point/logging"
	"github.com/sebdah/recharged/shared/websockets"
)

var (
	log      goLogging.Logger
	WsClient *websockets.Client
)

func main() {
	// Configure logging
	logging.Setup()

	log.Info("Starting re:charged charge-point simulator")
	log.Info("Environment: %s", config.Env)
	log.Info("Service port: %d", config.Config.GetInt("port"))

	// Parse command line options
	action := flag.String("action", "", "OCPP action")
	flag.Parse()

	// Connect to the WebSockets endpoint
	endpoint := fmt.Sprintf(
		"%s/%s",
		config.Config.GetString("central-system.endpoint-ocpp20j"),
		config.Config.GetString("identity"))
	log.Debug("Central-system endpoint: %s", endpoint)
	wsEndpoint, _ := url.Parse(endpoint)
	log.Debug("Connecting to %s over websockets", wsEndpoint.String())
	WsClient = websockets.NewClient(wsEndpoint)

	// Start the websockets communicator
	go websocketsCommunicator()
	go websocketPinger()

	// Send the actions
	if *action == "authorize" {
		WsClient.WriteChannel <- `[2, "1234", "Authorize", { "idTag": { "id": "1" } }]`
	} else if *action == "bootnotification" {
		WsClient.WriteChannel <- `[2, "1234", "BootNotification", { "chargePointModel": "Model X", "chargePointVendor": "Vendor Y"}]`
	} else if *action == "datatransfer" {
		WsClient.WriteChannel <- `[2, "1234", "DataTransfer", { "vendorId": "1234" }]`
	}

	// Do not terminate
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}

// Websockets message parser
func websocketsCommunicator() {
	var recv_msg string

	for {
		recv_msg = <-WsClient.ReadChannel

		log.Info("Received message: %s\n", recv_msg)
	}
}

// Send heartbeats via websocket pings
func websocketPinger() {
	WsClient.SendPing(config.Config.GetString("identity"))
	ticker := time.NewTicker(time.Duration(config.Config.GetInt("central-system.heartbeat-interval") * int(time.Second)))
	for {
		select {
		case <-ticker.C:
			WsClient.SendPing(config.Config.GetString("identity"))
		}
	}
}
