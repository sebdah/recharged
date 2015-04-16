package messages

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/sebdah/recharged/central-system/models"
	"github.com/sebdah/recharged/central-system/rpc"
	"github.com/sebdah/recharged/central-system/types"
	"gopkg.in/mgo.v2/bson"
)

type AuthorizeReq struct {
	IdTag types.IdToken `json:"idTag"`
}

type AuthorizeConf struct {
	IdTagInfo *types.IdTagInfo `json:"idTagInfo"`
	// PriceScheme, Not yet implemented
}

func NewAuthorizeReq() (req *AuthorizeReq) {
	req = new(AuthorizeReq)
	return
}

func NewAuthorizeConf() (conf *AuthorizeConf) {
	conf = new(AuthorizeConf)
	return
}

// Process the request, when it has been populated
func (this *AuthorizeReq) Process() (conf *AuthorizeConf, errorer rpc.Errorer) {
	// Get the IdTag
	idTag := models.NewIdTag()
	err := idTag.Collection().Find(bson.M{"idtag": this.IdTag.Id}).One(&idTag)
	if err != nil {
		if err.Error() != "not found" { // Not found documents are handled below
			log.Printf("MongoDB error: %s\n", err)
			errorer = rpc.NewPropertyConstraintViolation()
			errorer.SetDetails(`{"message": "IdTag not found"}`)
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
	} else if idTag.Id == "" { // The idTag does NOT exist
		idTagInfo.Status = types.AuthorizationStatusInvalid
	} else {
		idTagInfo.Status = types.AuthorizationStatusAccepted
	}

	// Populate the response configuration
	conf = new(AuthorizeConf)
	conf.IdTagInfo = idTagInfo
	//conf, err = json.Marshal(authorizeConf)
	//if err != nil {
	//	errorer = rpc.NewInternalError()
	//	errorer.SetDetails(fmt.Sprintf(`{"message": "%s"}`, err.Error()))
	//	return
	//}

	return
}

// Populate the object with JSON data
func (this *AuthorizeReq) ParseJson(data string) (err error) {
	decoder := json.NewDecoder(strings.NewReader(data))
	err = decoder.Decode(&this)
	if err != nil {
		log.Printf("Unable to parse payload: %s", err.Error())
		err = rpc.NewFormationViolation()
		return
	}

	return
}

// String representation
func (this *AuthorizeConf) String() (str string) {
	js, _ := json.Marshal(this)
	str = string(js)
	return
}
