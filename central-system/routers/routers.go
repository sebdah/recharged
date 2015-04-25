package routers

import (
	"github.com/gorilla/mux"
	"github.com/sebdah/recharged/central-system/req_processors"
	"github.com/sebdah/recharged/shared/transports"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// OCPP2.0-J
	wsServer := transports.NewWsServer()
	reqProcessor := req_processors.NewReqProcessor()
	wsServer.SetReqProcessor(reqProcessor)
	router.Path("/ocpp/v2.0j/ws").HandlerFunc(wsServer.Handler)

	return router
}
