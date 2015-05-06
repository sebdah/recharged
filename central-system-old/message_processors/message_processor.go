package message_processors

import "github.com/sebdah/recharged/admin/settings"

var configuration = settings.GetSettings()

type MessageProcessor struct{}

// Constructor
func NewMessageProcessor() (processor *MessageProcessor) {
	processor = new(MessageProcessor)
	return
}
