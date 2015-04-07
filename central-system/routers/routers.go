package routers

import (
	"github.com/gorilla/mux"
	"github.com/sebdah/recharged/central-system/handlers"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// Manager
	router.
		Path("/admin/idTags").
		Methods("GET").
		HandlerFunc(handlers.IdTagListHandler)
	router.
		Path("/admin/idTags").
		Methods("POST").
		HandlerFunc(handlers.IdTagCreateHandler)
	router.
		Path("/admin/idTags/{id}").
		Methods("GET").
		HandlerFunc(handlers.IdTagGetHandler)
	router.
		Path("/admin/idTags/{id}").
		Methods("PUT").
		HandlerFunc(handlers.IdTagUpdateHandler)

	// OCPP2.0-J
	router.
		Path("/ocpp/v2.0-j/authorize").
		Methods("POST").
		HandlerFunc(handlers.AuthorizeReqHandler)

	return router
}
