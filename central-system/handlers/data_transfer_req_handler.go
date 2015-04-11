package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sebdah/recharged/central-system/messages"
	"github.com/sebdah/recharged/central-system/types"
)

func DataTransferReqHandler(w http.ResponseWriter, r *http.Request) {
	// Populate the response configuration
	conf := new(messages.DataTransferConf)
	conf.Status = types.DataTransferStatusUnknownVendorId
	js, err := json.Marshal(conf)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	// Respond
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
