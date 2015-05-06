package rpc

import (
	"fmt"
	"regexp"
)

var callResultRegExp = regexp.MustCompile(`^\[(?P<messageId>\d+),(\ ?)"(?P<uniqueId>.+)",(\ ?)(?P<payload>.*)\]$`)

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
