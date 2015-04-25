package messages

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/sebdah/recharged/shared/rpc"
)

type DataTransferReq struct {
	messageType string `json:"-" type:"string"`
	VendorId    string `json:"vendorId" type:"string" required:"true"`
	MessageId   string `json:"messageId" type:"string" required:"false"`
	Data        string `json:"data" type:"string" required:"false"`
}

type DataTransferConf struct {
	Status string `json:"status" type:"string" required:"true"`
	Data   string `json:"data" type:"string" required:"false"`
}

func NewDataTransferReq() (req *DataTransferReq) {
	req = new(DataTransferReq)
	req.messageType = "DataTransfer"
	return
}

func NewDataTransferConf() (conf *DataTransferConf) {
	conf = new(DataTransferConf)
	return
}

// Get the message type
func (this *DataTransferReq) GetMessageType() string {
	return this.messageType
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
func (this *DataTransferReq) String() (str string) {
	js, _ := json.Marshal(this)
	str = string(js)
	return
}

// String representation
func (this *DataTransferConf) String() (str string) {
	js, _ := json.Marshal(this)
	str = string(js)
	return
}
