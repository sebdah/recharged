package routers

import (
	"github.com/gorilla/mux"
	"github.com/sebdah/recharged/admin/handlers"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// IdTags
	router.
		Path("/idtags").
		Methods("GET").
		HandlerFunc(handlers.IdTagListHandler)
	router.
		Path("/idtags").
		Methods("POST").
		HandlerFunc(handlers.IdTagCreateHandler)
	router.
		Path("/idtags/{id}").
		Methods("GET").
		HandlerFunc(handlers.IdTagGetHandler)
	router.
		Path("/idtags/{id}").
		Methods("DELETE").
		HandlerFunc(handlers.IdTagDeleteHandler)
	router.
		Path("/idtags/{id}").
		Methods("PUT").
		HandlerFunc(handlers.IdTagUpdateHandler)

	// Routing for charge points
	router.
		Path("/chargepoints").
		Methods("GET").
		HandlerFunc(handlers.ChargePointListHandler)
	router.
		Path("/chargepoints").
		Methods("POST").
		HandlerFunc(handlers.ChargePointCreateHandler)
	router.
		Path("/chargepoints/{id}").
		Methods("GET").
		HandlerFunc(handlers.ChargePointGetHandler)
	router.
		Path("/chargepoints/{id}").
		Methods("DELETE").
		HandlerFunc(handlers.ChargePointDeleteHandler)
	router.
		Path("/chargepoints/{id}").
		Methods("PUT").
		HandlerFunc(handlers.ChargePointUpdateHandler)
	router.
		Path("/chargepoints/validate/{vendor}/{model}").
		Methods("GET").
		HandlerFunc(handlers.ChargePointValidationHandler)

	return router
}
