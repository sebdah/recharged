package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/op/go-logging"
	"github.com/sebdah/recharged/admin/handlers"
)

var log logging.Logger

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
		Path("/chargepoints/validate").
		Methods("POST").
		HandlerFunc(handlers.ChargePointValidationHandler)

	// Routing for boot notification logs
	router.
		Path("/bootnotificationlogs").
		Methods("POST").
		HandlerFunc(handlers.BootNotificationLogCreateHandler)
	router.
		Path("/bootnotificationlogs/{id}").
		Methods("GET").
		HandlerFunc(handlers.BootNotificationLogGetHandler)

	return router
}

func HttpInterceptor(router http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		router.ServeHTTP(w, r)
		log.Debug("%s - %s %s", r.RemoteAddr, r.Method, r.URL)
	})
}
