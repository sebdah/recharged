package message_processors

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sebdah/recharged/central-system/models"
	"github.com/sebdah/recharged/shared/messages"
	"github.com/sebdah/recharged/shared/rpc"
	"github.com/sebdah/recharged/shared/types"
)

// Process the request
func (this *MessageProcessor) ProcessBootNotification(msg *messages.BootNotificationReq) (conf *messages.BootNotificationConf, errorer rpc.Errorer) {
	var err error

	// Build the ChargePoint
	chargePoint := types.NewChargePoint()
	chargePoint.SerialNumber = msg.ChargePointSerialNumber
	chargePoint.Vendor = msg.ChargePointVendor
	chargePoint.Model = msg.ChargePointModel
	chargePoint.Imsi = msg.Imsi

	// Ensure that the charge point exists in the system
	res, err := http.Get(fmt.Sprintf(
		"%s/chargepoints/validate/%s/%s",
		configuration.AdminServiceUrl.String(),
		chargePoint.Vendor,
		chargePoint.Model))
	if err != nil {
		conf = messages.NewBootNotificationConf(types.GenericStatusRejected)
		log.Printf("Error validating ChargePoint: %s", err.Error())
		return
	}
	if res.StatusCode == 404 {
		conf = messages.NewBootNotificationConf(types.GenericStatusRejected)
		return
	}

	// Update the database with the boot notification
	bootNotificationLog := models.NewBootNotificationLog()
	bootNotificationLog.Vendor = chargePoint.Vendor
	bootNotificationLog.Model = chargePoint.Model
	bootNotificationLog.SerialNumber = chargePoint.SerialNumber
	bootNotificationLog.Imsi = chargePoint.Imsi
	models.Save(bootNotificationLog)

	// Return success response
	conf = messages.NewBootNotificationConf(types.GenericStatusAccepted)

	return
}
