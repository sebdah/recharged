package message_processors

import (
	"github.com/sebdah/recharged/shared/messages"
	"github.com/sebdah/recharged/shared/rpc"
	"github.com/sebdah/recharged/shared/types"
)

func (this *MessageProcessor) ProcessDataTransferReq(msg *messages.DataTransferReq) (conf *messages.DataTransferConf, errorer rpc.Errorer) {
	// Populate the response configuration
	conf = messages.NewDataTransferConf()
	conf.Status = types.DataTransferStatusUnknownVendorId
	return
}
