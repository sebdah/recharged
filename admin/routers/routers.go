package routers

import (
	"github.com/gorilla/mux"
	"github.com/sebdah/recharged/admin/handlers"
)

func Router() *mux.Router {
	router := mux.NewRouter()

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
	adminRouter.
		Path("/chargePoints").
		Methods("GET").
		HandlerFunc(handlers.ChargePointListHandler)
	adminRouter.
		Path("/chargePoints").
		Methods("POST").
		HandlerFunc(handlers.ChargePointCreateHandler)
	adminRouter.
		Path("/chargePoints/{id}").
		Methods("GET").
		HandlerFunc(handlers.ChargePointGetHandler)
	adminRouter.
		Path("/chargePoints/{id}").
		Methods("DELETE").
		HandlerFunc(handlers.ChargePointDeleteHandler)
	adminRouter.
		Path("/chargePoints/{id}").
		Methods("PUT").
		HandlerFunc(handlers.ChargePointUpdateHandler)

	return router
}
