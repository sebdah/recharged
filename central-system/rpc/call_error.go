package rpc

import "fmt"

type CallError struct {
	UniqueId         string `type:"uniqueId" required:"true"`
	ErrorCode        string `type:"uniqueId" required:"true"`
	ErrorDescription string `type:"uniqueId" required:"true"`
	ErrorDetails     string `type:"uniqueId" required:"true"`
}

func NewCallError(uniqueId string, err Errorer) (callError *CallError) {
	callError = new(CallError)
	callError.UniqueId = uniqueId
	callError.ErrorCode = err.GetCode()
	callError.ErrorDescription = err.GetDescription()
	callError.ErrorDetails = err.GetDetails()
	return
}

// Get the string representation of the CallError
func (this *CallError) String() string {
	return fmt.Sprintf(
		`[4, "%s", "%s", %s]`,
		this.UniqueId,
		this.ErrorCode,
		this.ErrorDescription,
		this.ErrorDetails)
}
