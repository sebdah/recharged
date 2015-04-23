package routers

import (
	"github.com/gorilla/mux"
	"github.com/sebdah/recharged/central-system/transports"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// OCPP2.0-J
	router.Path("/ocpp/v2.0j/ws").HandlerFunc(transports.WebsocketTransport)

	return router
}
