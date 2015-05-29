package messages

import (
	"encoding/json"
	"strings"

	"github.com/sebdah/recharged/shared/rpc"
)

type DataTransferReq struct {
	messageType string `json:"-" type:"string"`
	VendorId    string `json:"vendorId" type:"string" required:"true"`
	MessageId   string `json:"messageId" type:"string" required:"false"`
	Data        string `json:"data" type:"string" required:"false"`

	*BaseReq
}

type DataTransferConf struct {
	Status string `json:"status" type:"string" required:"true"`
	Data   string `json:"data" type:"string" required:"false"`
}

// Constructor
func NewDataTransferReq(payload string) (req *DataTransferReq, rpcError rpc.Errorer) {
	req = new(DataTransferReq)
	req.messageType = "DataTransfer"

	decoder := json.NewDecoder(strings.NewReader(payload))
	err := decoder.Decode(&req)
	if err != nil {
		log.Notice("Unable to parse payload: %s", err.Error())
		rpcError = rpc.NewFormationViolation()
		return
	}

	return
}

// Constructor
func NewDataTransferConf(status string) (conf *DataTransferConf) {
	conf = new(DataTransferConf)
	conf.Status = status
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
