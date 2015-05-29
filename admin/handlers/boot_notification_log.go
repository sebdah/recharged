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

// Create new BootNotificationLog
func BootNotificationLogCreateHandler(w http.ResponseWriter, r *http.Request) {
	bootNotificationLog := models.NewBootNotificationLog()
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&bootNotificationLog)
	if err != nil {
		log.Debug("Unable to parse request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check required fields
	if bootNotificationLog.ChargePoint.Model == "" {
		log.Debug("Missing required field: chargePoint.model")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if bootNotificationLog.ChargePoint.Vendor == "" {
		log.Debug("Missing required field: chargePoint.vendor")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Save the object
	err = models.Save(bootNotificationLog)
	if err != nil {
		if mgo.IsDup(err) {
			w.WriteHeader(http.StatusConflict)
			log.Debug("BootNotificationLog already exists")
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error("Error in MongoDB communication: %s", err.Error())
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	js, _ := json.Marshal(bootNotificationLog)
	w.Write(js)

	return
}

func BootNotificationLogGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	collection := database.GetCollectionBootNotificationLogs()

	if bson.IsObjectIdHex(id) == false {
		http.NotFound(w, r)
		log.Debug("Invalid ObjectId '%s'", id)
		return
	}

	bootNotificationLog := new(models.BootNotificationLog)
	err := collection.FindId(bson.ObjectIdHex(id)).One(&bootNotificationLog)
	if err != nil {
		if strings.Contains(err.Error(), "not found") == true {
			http.NotFound(w, r)
			log.Debug("BootNotificationLog not found")
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Error("Error querying MongoDB: %s", err.Error())
			return
		}
	}

	data, err := json.Marshal(bootNotificationLog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Debug("MongoDB communication error: %s", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

	return
}
