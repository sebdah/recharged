package processors

type ReqProcessor struct{}

// Constructor
func NewReqProcessor() (processor *ReqProcessor) {
	processor = new(ReqProcessor)
	return
}
