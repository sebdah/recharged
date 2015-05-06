package rpc

import (
	"fmt"
	"regexp"
)

var callRegExp = regexp.MustCompile(`^\[(?P<messageId>\d+),(\ ?)"(?P<uniqueId>.+)",(\ ?)"(?P<action>\w+)",(\ ?)(?P<payload>.*)\]$`)

type Call struct {
	callCode int    `type:"int" required:"true"`
	UniqueId string `type:"string" required:"true"`
	Action   string `type:"string" required:"true"`
	Payload  string `type:"string" required:"true"`
}

// Constructor
func NewCall() (call *Call) {
	call = new(Call)
	call.callCode = 2
	call.UniqueId = ""
	call.Action = ""
	call.Payload = "{}"
	return
}

// Populate the variables with data from the request
func (this *Call) Populate(msg string) (callError *CallError) {
	match := callRegExp.FindStringSubmatch(msg)
	if len(match) == 0 {
		genericError := NewGenericError()
		genericError.SetDetails(`{"message": "Malformatted message"}`)
		callError = NewCallError(this.UniqueId, genericError)
		return
	}

	result := make(map[string]string)
	for i, name := range callRegExp.SubexpNames() {
		result[name] = match[i]
	}

	this.UniqueId = result["uniqueId"]
	this.Action = result["action"]
	this.Payload = result["payload"]

	// Convert payload to {} if it's set to "null"
	if this.Payload == "null" {
		this.Payload = "{}"
	}

	return
}

// Get the string representation
func (this *Call) String() string {
	return fmt.Sprintf(
		`[%d, "%s", "%s", %s]`,
		this.callCode,
		this.UniqueId,
		this.Action,
		this.Payload)
}
