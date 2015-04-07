package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sebdah/recharged/central-system/database"
	"github.com/sebdah/recharged/central-system/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func IdTagCreateHandler(w http.ResponseWriter, r *http.Request) {
	idTag := models.NewIdTag()
	decoder := json.NewDecoder(r.Body)
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

func IdTagGetHandler(w http.ResponseWriter, r *http.Request) {
	var idTag = new(models.IdTag)

	vars := mux.Vars(r)
	id := vars["id"]
	collection := database.GetCollectionIdTags()

	err := collection.Find(bson.M{"idtag": id}).One(&idTag)
	if err != nil {
		if strings.Contains(err.Error(), "not found") == true {
			http.NotFound(w, r)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("Error querying MongoDB: %s", err)
			return
		}
	}

	data, err := json.Marshal(idTag)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error marshalling JSON: %s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

	return
}

func IdTagListHandler(w http.ResponseWriter, r *http.Request) {
	var idTags []models.IdTag
	collection := database.GetCollectionIdTags()
	err := collection.Find(bson.M{}).All(&idTags)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error querying MongoDB: %s", err)
		return
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

func IdTagUpdateHandler(w http.ResponseWriter, r *http.Request) {
	var idTag = new(models.IdTag)

	vars := mux.Vars(r)
	id := vars["id"]
	collection := database.GetCollectionIdTags()

	err := collection.Find(bson.M{"idtag": id}).One(&idTag)
	if err != nil {
		if strings.Contains(err.Error(), "not found") == true {
			http.NotFound(w, r)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("Error querying MongoDB: %s", err)
			return
		}
	}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&idTag)
	if err != nil {
		log.Printf("Unable to parse request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = models.Update(idTag)
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
