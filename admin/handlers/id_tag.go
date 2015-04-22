package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sebdah/recharged/admin/database"
	"github.com/sebdah/recharged/admin/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Get an IdTag from the Database
func getIdTag(w http.ResponseWriter, r *http.Request) (idTag models.IdTag) {
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
			log.Printf("Error querying MongoDB: %s\n", err.Error())
			return
		}
	}

	return
}

// Create new IdTag
func IdTagCreateHandler(w http.ResponseWriter, r *http.Request) {
	idTag := models.NewIdTag()
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&idTag)
	if err != nil {
		log.Printf("Unable to parse request: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check that the idTag field is set
	if idTag.IdTag == "" {
		log.Printf("Missing idTag in request")
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

// Delete IdTag
func IdTagDeleteHandler(w http.ResponseWriter, r *http.Request) {
	idTag := getIdTag(w, r)
	err := models.Delete(&idTag)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error deleting IdTag %s", idTag.Id)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

// Get an IdTag
func IdTagGetHandler(w http.ResponseWriter, r *http.Request) {
	idTag := getIdTag(w, r)

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

// List all IdTags
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

// Update an IdTag
func IdTagUpdateHandler(w http.ResponseWriter, r *http.Request) {
	idTag := getIdTag(w, r)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&idTag)
	if err != nil {
		log.Printf("Unable to parse request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = models.Update(&idTag)
	if err != nil {
		if mgo.IsDup(err) {
			w.WriteHeader(http.StatusConflict)
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)

	return
}
