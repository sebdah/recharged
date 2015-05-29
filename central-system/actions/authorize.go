package actions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sebdah/recharged/central-system/config"
	"github.com/sebdah/recharged/shared/messages"
	"github.com/sebdah/recharged/shared/rpc"
	"github.com/sebdah/recharged/shared/types"
)

// Handle Authorize requests
func Authorize(req *messages.AuthorizeReq) (conf *messages.AuthorizeConf, errorer rpc.Errorer) {
	// Get the IdTag
	idTag := types.NewIdTag()
	res, err := http.Get(config.Config.GetString("admin-service.endpoint") + "/idtags/" + req.IdTag.Id)
	if err != nil {
		log.Error(err.Error())
		errorer = rpc.NewInternalError()
		errorer.SetDetails(fmt.Sprintf(`{"message": "%s"}`, err.Error()))
		return
	}

	// Handle 404s
	if res.StatusCode == 404 {
		log.Error("IdTag %d not found", req.IdTag.Id)
		errorer = rpc.NewPropertyConstraintViolation()
		errorer.SetDetails(`{"message": "idTag not found"}`)
		return
	}

	// Populate the response
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&idTag)
	if err != nil {
		log.Error(err.Error())
		errorer = rpc.NewInternalError()
		errorer.SetDetails(fmt.Sprintf(`{"message": "%s"}`, err.Error()))
		return
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
	} else if idTag.IdTag == "" { // The idTag does NOT exist
		idTagInfo.Status = types.AuthorizationStatusInvalid
	} else {
		idTagInfo.Status = types.AuthorizationStatusAccepted
	}
	log.Debug("IdTag: %+v", idTag)

	// Populate the response configuration
	conf = messages.NewAuthorizeConf()
	conf.IdTagInfo = idTagInfo

	return
}
