package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sebdah/recharged/central-system/models"
	"gopkg.in/mgo.v2"
)

func IdTagAddHandler(rw http.ResponseWriter, req *http.Request) {
	idTag := models.NewIdTag()
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&idTag)
	if err != nil {
		log.Printf("Unable to parse request")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err = models.Save(idTag)
	if err != nil {
		if mgo.IsDup(err) {
			rw.WriteHeader(http.StatusConflict)
			return
		} else {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	rw.WriteHeader(http.StatusCreated)

	return
}
