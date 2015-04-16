package routers

import (
	"github.com/gorilla/mux"
	"github.com/sebdah/recharged/central-system/handlers"
	"github.com/sebdah/recharged/central-system/transports"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// OCPP2.0-J
	router.Path("/ocpp/v2.0j/ws").HandlerFunc(transports.WebsocketTransport)

	// Manager
	adminRouter := router.PathPrefix("/admin").Subrouter()
	adminRouter.
		Path("/idTags").
		Methods("GET").
		HandlerFunc(handlers.IdTagListHandler)
	adminRouter.
		Path("/idTags").
		Methods("POST").
		HandlerFunc(handlers.IdTagCreateHandler)
	adminRouter.
		Path("/idTags/{id}").
		Methods("GET").
		HandlerFunc(handlers.IdTagGetHandler)
	adminRouter.
		Path("/idTags/{id}").
		Methods("DELETE").
		HandlerFunc(handlers.IdTagDeleteHandler)
	adminRouter.
		Path("/idTags/{id}").
		Methods("PUT").
		HandlerFunc(handlers.IdTagUpdateHandler)

	return router
}
