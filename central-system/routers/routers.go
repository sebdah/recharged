package routers

import (
	"github.com/gorilla/mux"
	"github.com/sebdah/recharged/central-system/handlers"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	idtag := router.Path("/idtag/{id}").Subrouter()
	idtag.Path("/authorize").Methods("POST").HandlerFunc(handlers.AuthorizeHandler)

	return router
}
