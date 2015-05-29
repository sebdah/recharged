package heartbeat

import "github.com/op/go-logging"

var (
	HeartbeatChannel chan string
	log              logging.Logger
)

func init() {
	HeartbeatChannel = make(chan string)
	go handler()
}

func handler() {
	log.Debug("Starting HeartbeatHandler")

	for {
		message := <-HeartbeatChannel
		log.Debug("Received heartbeat from %s", message)
	}
}
