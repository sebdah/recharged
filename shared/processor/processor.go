package processor

import (
	"github.com/sebdah/recharged/shared/messages"
	"github.com/sebdah/recharged/shared/rpc"
)

type Processor interface {
	ProcessAuthorize(msg *messages.AuthorizeReq) (conf *messages.AuthorizeConf, errorer rpc.Errorer)
	ProcessBootNotification(msg *messages.BootNotificationReq) (conf *messages.BootNotificationConf, errorer rpc.Errorer)
	ProcessDataTransfer(msg *messages.DataTransferReq) (conf *messages.DataTransferConf, errorer rpc.Errorer)
}
