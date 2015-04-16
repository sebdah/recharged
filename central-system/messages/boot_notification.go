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

func NewBootNotificationConf() (conf *BootNotificationConf) {
	conf = new(BootNotificationConf)
	return
}

// Process the request
func (this *BootNotificationReq) Process() (conf *BootNotificationConf, errorer rpc.Errorer) {
	var err error

	// Build the ChargePoint
	chargePoint := models.NewChargePoint()
	chargePoint.SerialNumber = this.ChargePointSerialNumber
	chargePoint.Vendor = this.ChargePointVendor
	chargePoint.Model = this.ChargePointModel
	chargePoint.Imsi = this.Imsi

	// Set status and update database
	status := types.GenericStatusNotSupported
	if this.ChargePointVendor == "" { // Vendor is required
		status = types.GenericStatusRejected
	} else if this.ChargePointModel == "" { // Model is required
		status = types.GenericStatusRejected
	} else if this.ChargePointSerialNumber != "" { // serialNumber:set
		status = types.GenericStatusAccepted
		err = models.Upsert(
			bson.M{"serialnumber": chargePoint.SerialNumber},
			chargePoint)
	} else if this.Imsi != "" {
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
	conf = new(BootNotificationConf)
	conf.CurrentTime.Time = time.Now()
	conf.HeartbeatInterval = 10
	conf.Status = status

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
func (this *BootNotificationConf) String() (str string) {
	js, _ := json.Marshal(this)
	str = string(js)
	return
}
