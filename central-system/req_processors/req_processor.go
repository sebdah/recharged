package req_processors

import "github.com/sebdah/recharged/admin/settings"

var configuration = settings.GetSettings()

type ReqProcessor struct{}

// Constructor
func NewReqProcessor() (processor *ReqProcessor) {
	processor = new(ReqProcessor)
	return
}
