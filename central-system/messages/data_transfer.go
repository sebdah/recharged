package messages

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/sebdah/recharged/central-system/rpc"
	"github.com/sebdah/recharged/central-system/types"
)

type DataTransferReq struct {
	VendorId  string `json:"vendorId" type:"string" required:"true"`
	MessageId string `json:"messageId" type:"string" required:"false"`
	Data      string `json:"data" type:"string" required:"false"`
}

type DataTransferConf struct {
	Status string `json:"status" type:"string" required:"true"`
	Data   string `json:"data" type:"string" required:"false"`
}

func NewDataTransferReq() (req *DataTransferReq) {
	req = new(DataTransferReq)
	return
}

func NewDataTransferConf() (conf *DataTransferConf) {
	conf = new(DataTransferConf)
	return
}

// Process
func (this *DataTransferReq) Process() (conf *DataTransferConf, errorer rpc.Errorer) {
	// Populate the response configuration
	conf = NewDataTransferConf()
	conf.Status = types.DataTransferStatusUnknownVendorId
	return
}

// Populate the object with JSON data
func (this *DataTransferReq) ParseJson(data string) (err error) {
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
func (this *DataTransferConf) String() (str string) {
	js, _ := json.Marshal(this)
	str = string(js)
	return
}
