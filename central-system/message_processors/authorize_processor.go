package message_processors

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sebdah/recharged/shared/messages"
	"github.com/sebdah/recharged/shared/rpc"
	"github.com/sebdah/recharged/shared/types"
)

// Process the request, when it has been populated
func (this *MessageProcessor) ProcessAuthorizeReq(msg *messages.AuthorizeReq) (conf *messages.AuthorizeConf, errorer rpc.Errorer) {
	// Get the IdTag
	idTag := types.NewIdTag()
	res, err := http.Get(configuration.AdminServiceUrl.String() + "/idtags/" + msg.IdTag.Id)
	if err != nil {
		errorer = rpc.NewInternalError()
		errorer.SetDetails(fmt.Sprintf(`{"message": "%s"}`, err.Error()))
		return
	}
	if res.StatusCode == 404 {
		errorer = rpc.NewPropertyConstraintViolation()
		errorer.SetDetails(`{"message": "IdTag not found"}`)
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
	} else if idTag.Id == "" { // The idTag does NOT exist
		idTagInfo.Status = types.AuthorizationStatusInvalid
	} else {
		idTagInfo.Status = types.AuthorizationStatusAccepted
	}

	// Populate the response configuration
	conf = messages.NewAuthorizeConf()
	conf.IdTagInfo = idTagInfo

	return
}
