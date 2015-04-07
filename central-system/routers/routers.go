package routers

import (
	"github.com/gorilla/mux"
	"github.com/sebdah/recharged/central-system/handlers"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// Manager
	adminRouter := router.Path("/admin").Subrouter()
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
		Methods("PUT").
		HandlerFunc(handlers.IdTagUpdateHandler)

	// OCPP2.0-J
	occp20jRouter := router.Path("/ocpp/v2.0j").Subrouter()
	occp20jRouter.
		Path("/authorize").
		Methods("POST").
		HandlerFunc(handlers.AuthorizeReqHandler)

	return router
}
