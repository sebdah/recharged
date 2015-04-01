package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sebdah/recharged/central-system/messages"
	"github.com/sebdah/recharged/central-system/types"
)

func AuthorizeReqHandler(rw http.ResponseWriter, req *http.Request) {
	authorizeReq := new(messages.AuthorizeReq)

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&authorizeReq)
	if err != nil {
		log.Printf("Unable to parse Authorize.req")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	idTagInfo := new(types.IdTagInfo)
	idTagInfo.Status = types.AuthorizationStatusAccepted

	authorizeConf := new(messages.AuthorizeConf)
	authorizeConf.IdTagInfo = idTagInfo

	authConfJson, err := json.Marshal(authorizeConf)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(authConfJson)
}
