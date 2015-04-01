package routers

import (
	"github.com/gorilla/mux"
	"github.com/sebdah/recharged/central-system/handlers"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.
		Path("/ocpp/v2.0-j/authorize").
		Methods("POST").
		HandlerFunc(handlers.AuthorizeReqHandler)

	return router
}
