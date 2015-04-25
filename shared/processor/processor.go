package processor

import (
	"github.com/sebdah/recharged/shared/messages"
	"github.com/sebdah/recharged/shared/rpc"
)

type Processor interface {
	ProcessAuthorizeReq(msg *messages.AuthorizeReq) (conf *messages.AuthorizeConf, errorer rpc.Errorer)
	ProcessBootNotificationReq(msg *messages.BootNotificationReq) (conf *messages.BootNotificationConf, errorer rpc.Errorer)
	ProcessDataTransferReq(msg *messages.DataTransferReq) (conf *messages.DataTransferConf, errorer rpc.Errorer)
}
