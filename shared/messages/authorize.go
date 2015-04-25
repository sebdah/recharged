package messages

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/sebdah/recharged/shared/rpc"
	"github.com/sebdah/recharged/shared/types"
)

type AuthorizeReq struct {
	messageType string        `json:"-" type:"string"`
	IdTag       types.IdToken `json:"idTag"`
}

type AuthorizeConf struct {
	IdTagInfo *types.IdTagInfo `json:"idTagInfo"`
	// PriceScheme, Not yet implemented
}

func NewAuthorizeReq() (req *AuthorizeReq) {
	req = new(AuthorizeReq)
	req.messageType = "Authorize"
	return
}

func NewAuthorizeConf() (conf *AuthorizeConf) {
	conf = new(AuthorizeConf)
	return
}

// Get the message type
func (this *AuthorizeReq) GetMessageType() string {
	return this.messageType
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
func (this *AuthorizeReq) String() (str string) {
	js, _ := json.Marshal(this)
	str = string(js)
	return
}

// String representation
func (this *AuthorizeConf) String() (str string) {
	js, _ := json.Marshal(this)
	str = string(js)
	return
}
