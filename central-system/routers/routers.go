package routers

import (
	"github.com/gorilla/mux"
	"github.com/sebdah/recharged/central-system/handlers"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.
		Path("/idtag/{id}/authorize").
		Methods("GET").
		HandlerFunc(handlers.AuthorizeHandler)

	return router
}
