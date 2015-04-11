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

// Handle the incoming BootNotification.Req
func BootNotificationReqHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the incoming request
	req := new(messages.BootNotificationReq)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		log.Printf("Unable to parse BootNotification.req: %s\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Build the ChargePoint
	chargePoint := models.NewChargePoint()
	chargePoint.SerialNumber = req.ChargePointSerialNumber
	chargePoint.Vendor = req.ChargePointVendor
	chargePoint.Model = req.ChargePointModel
	chargePoint.Imsi = req.Imsi

	// Set status and update database
	status := types.GenericStatusNotSupported
	if req.ChargePointVendor == "" { // Vendor is required
		status = types.GenericStatusRejected
	} else if req.ChargePointModel == "" { // Model is required
		status = types.GenericStatusRejected
	} else if req.ChargePointSerialNumber != "" { // serialNumber:set
		status = types.GenericStatusAccepted
		err = models.Upsert(
			bson.M{"serialnumber": chargePoint.SerialNumber},
			chargePoint)
	} else if req.Imsi != "" {
		status = types.GenericStatusAccepted
		err = models.Upsert(
			bson.M{"imsi": chargePoint.Imsi},
			chargePoint)
	} else {
		status = types.GenericStatusAccepted
		err = models.Upsert(
			bson.M{
				"vendor": chargePoint.Vendor,
				"model":  chargePoint.Model,
			},
			chargePoint)
	}
	if err != nil {
		log.Printf("Error validating ChargePoint: %s", err)
		status = types.GenericStatusRejected
	}

	// Build response
	conf := new(messages.BootNotificationConf)
	conf.CurrentTime.Time = time.Now()
	conf.HeartbeatInterval = 10
	conf.Status = status

	// Marshal JSON
	js, err := json.Marshal(conf)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error marshalling JSON: %s\n", err)
		return
	}

	// Respond
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
