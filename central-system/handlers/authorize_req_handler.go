package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/sebdah/recharged/central-system/messages"
	"github.com/sebdah/recharged/central-system/models"
	"github.com/sebdah/recharged/central-system/types"
	"gopkg.in/mgo.v2/bson"
)

func AuthorizeReqHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the incoming request
	authorizeReq := new(messages.AuthorizeReq)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&authorizeReq)
	if err != nil {
		log.Printf("Unable to parse Authorize.req: %s\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get the IdTag
	idTag := models.IdTag{}
	err = idTag.Collection().Find(bson.M{"idtag": authorizeReq.IdTag.Id}).One(&idTag)
	if err != nil {
		if err.Error() == "not found" {
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			log.Printf("MongoDB error: %s\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Create the IdTagInfo for the response
	idTagInfo := types.NewIdTagInfo()

	// Set the status flag
	beginning, _ := time.Parse(time.RFC3339, "0001-01-01T00:00:00Z")
	if idTag.Active == false { // Check for deactivation
		idTagInfo.Status = types.AuthorizationStatusBlocked
	} else if idTag.ExpiryDate.Equal(beginning) == false {
		if idTag.ExpiryDate.Before(time.Now().UTC()) == true {
			idTagInfo.Status = types.AuthorizationStatusExpired
		}
	} else {
		idTagInfo.Status = types.AuthorizationStatusAccepted
	}

	// Populate the response configuration
	authorizeConf := new(messages.AuthorizeConf)
	authorizeConf.IdTagInfo = idTagInfo
	authConfJson, err := json.Marshal(authorizeConf)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	// Respond
	w.WriteHeader(http.StatusOK)
	w.Write(authConfJson)
}
