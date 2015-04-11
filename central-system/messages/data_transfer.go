package messages

type DataTransferReq struct {
	VendorId  string `json:"vendorId" type:"string" required:"true"`
	MessageId string `json:"messageId" type:"string" required:"false"`
	Data      string `json:"data" type:"string" required:"false"`
}

type DataTransferConf struct {
	Status string `json:"status" type:"string" required:"true"`
	Data   string `json:"data" type:"string" required:"false"`
}
