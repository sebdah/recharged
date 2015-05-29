package actions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/sebdah/recharged/central-system/config"
	"github.com/sebdah/recharged/shared/messages"
	"github.com/sebdah/recharged/shared/rpc"
	"github.com/sebdah/recharged/shared/types"
)

func BootNotification(req messages.BootNotificationReq) (conf *messages.BootNotificationConf, errorer rpc.Errorer) {
	var err error

	// Build the ChargePoint
	chargePoint := types.NewChargePoint()
	chargePoint.Id = req.ChargePoint
	chargePoint.SerialNumber = req.ChargePointSerialNumber
	chargePoint.Vendor = req.ChargePointVendor
	chargePoint.Model = req.ChargePointModel
	chargePoint.Imsi = req.Imsi

	// Ensure that the charge point exists in the system
	url := fmt.Sprintf("%s/chargepoints/validate", config.Config.GetString("admin-service.endpoint"))
	js, _ := json.Marshal(chargePoint)
	res, err := http.Post(url, "", strings.NewReader(string(js)))
	if err != nil {
		conf = messages.NewBootNotificationConf(types.GenericStatusRejected)
		log.Error("Error validating ChargePoint: %s", err.Error())
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		conf = messages.NewBootNotificationConf(types.GenericStatusRejected)
		log.Debug("Unexpected response status code %d from %s", res.StatusCode, url)
		return
	}

	// Create BootNotificationLog
	bootNotificationLog := types.NewBootNotificationLog()
	bootNotificationLog.ChargePoint = *chargePoint
	js, _ = json.Marshal(bootNotificationLog)

	// Update the database with the boot notification
	url = fmt.Sprintf("%s/bootnotificationlogs", config.Config.GetString("admin-service.endpoint"))
	res, err = http.Post(url, "", strings.NewReader(string(js)))
	if err != nil {
		conf = messages.NewBootNotificationConf(types.GenericStatusRejected)
		log.Error("Error creating boot notification log: %s", err.Error())
		return
	}
	defer res.Body.Close()

	// Return success response
	conf = messages.NewBootNotificationConf(types.GenericStatusAccepted)

	return
}
