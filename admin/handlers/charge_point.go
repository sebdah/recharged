package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sebdah/recharged/admin/database"
	"github.com/sebdah/recharged/admin/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Fetch ChargePoint from the database
func getChargePoint(w http.ResponseWriter, r *http.Request) (chargePoint models.ChargePoint) {
	vars := mux.Vars(r)
	id := vars["id"]
	collection := database.GetCollectionChargePoints()

	err := collection.Find(bson.M{"id": id}).One(&chargePoint)
	if err != nil {
		if strings.Contains(err.Error(), "not found") == true {
			http.NotFound(w, r)
			log.Debug("Charge point '%s' not found", id)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Error("Error querying MongoDB: %s", err.Error())
			return
		}
	}

	return
}

// Create new ChargePoint
func ChargePointCreateHandler(w http.ResponseWriter, r *http.Request) {
	chargePoint := models.NewChargePoint()
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&chargePoint)
	if err != nil {
		log.Debug("Unable to parse request: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check that the id field is set
	if chargePoint.ExternalId == "" {
		log.Debug("Missing id in request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check that the model field is set
	if chargePoint.Model == "" {
		log.Debug("Missing model in request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check that the vendor field is set
	if chargePoint.Vendor == "" {
		log.Debug("Missing vendor in request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = models.Save(chargePoint)
	if err != nil {
		if mgo.IsDup(err) {
			w.WriteHeader(http.StatusConflict)
			log.Debug("ChargePoint already exists")
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error("Error in MongoDB communication: %s", err.Error())
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	js, _ := json.Marshal(chargePoint)
	w.Write(js)

	return
}

// Delete ChargePoint
func ChargePointDeleteHandler(w http.ResponseWriter, r *http.Request) {
	chargePoint := getChargePoint(w, r)
	err := models.Delete(&chargePoint)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error(
			"Error deleting ChargePoint '%s'. Reason: %s",
			chargePoint.ExternalId,
			err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

// Get a ChargePoint
func ChargePointGetHandler(w http.ResponseWriter, r *http.Request) {
	chargePoint := getChargePoint(w, r)

	data, err := json.Marshal(chargePoint)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Debug("Error marshalling JSON: %s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

	return
}

// List all ChargePoints
func ChargePointListHandler(w http.ResponseWriter, r *http.Request) {
	var chargePoints []models.ChargePoint
	collection := database.GetCollectionChargePoints()
	err := collection.Find(bson.M{}).All(&chargePoints)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error("Error querying MongoDB: %s", err)
		return
	}

	data, err := json.Marshal(chargePoints)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Debug("Error marshalling JSON: %s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

	return
}

// Update a ChargePoint
func ChargePointUpdateHandler(w http.ResponseWriter, r *http.Request) {
	chargePoint := getChargePoint(w, r)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&chargePoint)
	if err != nil {
		log.Debug("Unable to parse request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = models.Update(&chargePoint)
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

// Validate if a model/vendor combination exists
// Returns 200 if found, else 404
func ChargePointValidationHandler(w http.ResponseWriter, r *http.Request) {
	chargePoint := models.NewChargePoint()
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&chargePoint)
	if err != nil {
		log.Debug("Unable to parse request: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	collection := database.GetCollectionChargePoints()
	cnt, err := collection.Find(bson.M{"id": chargePoint.ExternalId}).Count()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error("Error querying MongoDB: %s", err.Error())
		return
	}

	if cnt >= 1 {
		log.Debug("Found ChargePoint with id %s", chargePoint.ExternalId)
		w.WriteHeader(http.StatusOK)
		return
	}

	log.Debug("No ChargePoint with id %s", chargePoint.ExternalId)
	w.WriteHeader(http.StatusNotFound)
	return
}
