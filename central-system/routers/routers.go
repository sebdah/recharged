package routers

import (
	"github.com/gorilla/mux"
	"github.com/sebdah/recharged/central-system/processors"
	"github.com/sebdah/recharged/shared/transports"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// Create websockets server
	wsServer := transports.NewWsServer()

	// Configure request processors
	reqProcessor := processors.NewReqProcessor()
	confProcessor := processors.NewConfProcessor()
	wsServer.SetReqProcessor(reqProcessor)
	wsServer.SetConfProcessor(confProcessor)

	// OCPP2.0-J
	router.Path("/ocpp/v2.0j/ws").HandlerFunc(wsServer.Handler)

	return router
}
