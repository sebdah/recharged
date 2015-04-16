package rpc

import "regexp"

var callRegExp = regexp.MustCompile(`^\[(?P<messageId>\d+),(\ ?)"(?P<uniqueId>.+)",(\ ?)"(?P<action>\w+)",(\ ?)(?P<payload>.*)\]$`)

type Call struct {
	UniqueId string `type:"string" required:"true"`
	Action   string `type:"string" required:"true"`
	Payload  string `type:"string" required:"true"`
}

// Constructor
func NewCall() (call *Call) {
	call = new(Call)
	call.UniqueId = ""
	call.Action = ""
	call.Payload = "{}"
	return
}

// Populate the variables with data from the request
func (this *Call) Populate(msg string) (callError *CallError) {
	match := callRegExp.FindStringSubmatch(msg)
	result := make(map[string]string)
	for i, name := range callRegExp.SubexpNames() {
		result[name] = match[i]
	}

	this.UniqueId = result["uniqueId"]
	this.Action = result["action"]
	this.Payload = result["payload"]

	// Check for missing UniqueId
	if this.UniqueId == "" {
		callError = NewCallError(this.UniqueId, NewGenericError())
	}

	// Check for missing Action
	if this.Action == "" {
		callError = NewCallError(this.UniqueId, NewNotImplementedError())
	}

	// Convert payload to {} if it's set to "null"
	if this.Payload == "null" {
		this.Payload = "{}"
	}

	return
}
