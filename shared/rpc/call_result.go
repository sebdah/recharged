package rpc

import (
	"errors"
	"fmt"
	"regexp"
)

type CallResult struct {
	UniqueId string `type:"string" required:"true"`
	Payload  string `type:"string" required:"true"`
}

// Constructor
func NewCallResult(uniqueId string) (callResult *CallResult) {
	callResult = new(CallResult)
	callResult.UniqueId = uniqueId
	callResult.Payload = "{}"
	return
}

// Populate the variables with data from the request
func (this *CallResult) Populate(msg string) (err error) {
	callResultRegExp := regexp.MustCompile(`^\[(?P<messageId>\d+),(\ ?)"(?P<uniqueId>.+)",(\ ?)(?P<payload>.*)\]$`)

	match := callResultRegExp.FindStringSubmatch(msg)
	if len(match) == 0 {
		err = errors.New(fmt.Sprintf("CallResult '%s' does not match regular expression", msg))
	}

	result := make(map[string]string)
	for i, name := range callResultRegExp.SubexpNames() {
		result[name] = match[i]
	}

	this.UniqueId = result["uniqueId"]
	this.Payload = result["payload"]

	// Convert payload to {} if it's set to "null"
	if this.Payload == "null" {
		this.Payload = "{}"
	}

	return
}

// Set payload
func (this *CallResult) SetPayload(payload string) {
	this.Payload = payload
}

// Get the string representation
func (this *CallResult) String() string {
	return fmt.Sprintf(
		`[3, "%s", %s]`,
		this.UniqueId,
		this.Payload)
}
