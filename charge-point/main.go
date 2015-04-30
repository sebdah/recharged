package main

import "github.com/sebdah/recharged/shared/transports"

func main() {
	wsClient := transports.NewWsClient()
	conn := wsClient.Connect()
}
