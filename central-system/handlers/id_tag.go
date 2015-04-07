package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sebdah/recharged/central-system/database"
	"github.com/sebdah/recharged/central-system/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func IdTagCreateHandler(w http.ResponseWriter, req *http.Request) {
	idTag := models.NewIdTag()
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&idTag)
	if err != nil {
		log.Printf("Unable to parse request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = models.Save(idTag)
	if err != nil {
		if mgo.IsDup(err) {
			w.WriteHeader(http.StatusConflict)
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)

	return
}

func IdTagListHandler(w http.ResponseWriter, req *http.Request) {
	var idTags []models.IdTag
	collection := database.GetCollectionIdTags()
	err := collection.Find(bson.M{}).All(&idTags)
	if err != nil {
		log.Fatal(err)
	}

	data, err := json.Marshal(idTags)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error marshalling JSON: %s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

	return
}
