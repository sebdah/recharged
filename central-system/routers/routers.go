package routers

import (
	"github.com/gorilla/mux"
	"github.com/sebdah/recharged/central-system/message_processors"
	"github.com/sebdah/recharged/shared/transports"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// OCPP2.0-J
	wsServer := transports.NewWsServer()
	processor := message_processors.NewMessageProcessor()
	wsServer.SetProcessor(processor)
	router.Path("/ocpp/v2.0j/ws").HandlerFunc(wsServer.Handler)

	return router
}
