package processors

type ConfProcessor struct{}

// Constructor
func NewConfProcessor() (processor *ConfProcessor) {
	processor = new(ConfProcessor)
	return
}
