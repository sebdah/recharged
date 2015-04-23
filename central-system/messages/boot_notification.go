package messages

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/sebdah/recharged/central-system/models"
	"github.com/sebdah/recharged/central-system/rpc"
	"github.com/sebdah/recharged/central-system/types"
)

type BootNotificationReq struct {
	ChargePointModel        string `json:"chargePointModel" type:"string" required:"true"`
	ChargePointVendor       string `json:"chargePointVendor" type:"string" required:"true"`
	ChargePointSerialNumber string `json:"chargePointSerialNumber" type:"string" required:"false"`
	Imsi                    string `json:"imsi" type:"string" required:"false"`
}

type BootNotificationConf struct {
	CurrentTime       types.JSONTime `json:"currentTime" type:"time.Time" required:"true"`
	HeartbeatInterval int64          `json:"heartbeatInterval" type:"int" required:"false"`
	Status            string         `json:"status" type:"string" required:"true"`
}

func NewBootNotificationReq() (req *BootNotificationReq) {
	req = new(BootNotificationReq)
	return
}

func NewBootNotificationConf(status string) (conf *BootNotificationConf) {
	conf = new(BootNotificationConf)
	conf.CurrentTime.Time = time.Now()
	conf.HeartbeatInterval = 10
	conf.Status = status
	return
}

// Process the request
func (this *BootNotificationReq) Process() (conf *BootNotificationConf, errorer rpc.Errorer) {
	var err error

	// Build the ChargePoint
	chargePoint := types.NewChargePoint()
	chargePoint.SerialNumber = this.ChargePointSerialNumber
	chargePoint.Vendor = this.ChargePointVendor
	chargePoint.Model = this.ChargePointModel
	chargePoint.Imsi = this.Imsi

	// Ensure that the charge point exists in the system
	res, err := http.Get(fmt.Sprintf(
		"%s/chargepoints/validate/%s/%s",
		configuration.AdminServiceUrl.String(),
		chargePoint.Vendor,
		chargePoint.Model))
	if err != nil {
		conf = NewBootNotificationConf(types.GenericStatusRejected)
		log.Printf("Error validating ChargePoint: %s", err.Error())
		return
	}
	if res.StatusCode == 404 {
		conf = NewBootNotificationConf(types.GenericStatusRejected)
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
	conf = NewBootNotificationConf(types.GenericStatusAccepted)

	return
}

// Populate the object with JSON data
func (this *BootNotificationReq) ParseJson(data string) (err error) {
	decoder := json.NewDecoder(strings.NewReader(data))
	err = decoder.Decode(&this)
	if err != nil {
		log.Printf("Unable to parse payload: %s", err.Error())
		return
	}

	return
}

// String representation
func (this *BootNotificationReq) String() (str string) {
	js, _ := json.Marshal(this)
	str = string(js)
	return
}

// String representation
func (this *BootNotificationConf) String() (str string) {
	js, _ := json.Marshal(this)
	str = string(js)
	return
}
