package processors

import (
	"github.com/sebdah/recharged/shared/messages"
	"github.com/sebdah/recharged/shared/rpc"
	"github.com/sebdah/recharged/shared/types"
)

func (this *ReqProcessor) ProcessDataTransferReq(msg *messages.DataTransferReq) (conf *messages.DataTransferConf, errorer rpc.Errorer) {
	// Populate the response configuration
	conf = messages.NewDataTransferConf()
	conf.Status = types.DataTransferStatusUnknownVendorId
	return
}
