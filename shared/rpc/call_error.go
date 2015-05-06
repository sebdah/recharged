package rpc

import "fmt"

type CallError struct {
	callCode         int    `type:"int" required:"true"`
	UniqueId         string `type:"uniqueId" required:"true"`
	ErrorCode        string `type:"uniqueId" required:"true"`
	ErrorDescription string `type:"uniqueId" required:"true"`
	ErrorDetails     string `type:"uniqueId" required:"true"`
}

func NewCallError(uniqueId string, err Errorer) (callError *CallError) {
	callError = new(CallError)
	callError.callCode = 4
	callError.UniqueId = uniqueId
	callError.ErrorCode = err.GetCode()
	callError.ErrorDescription = err.GetDescription()
	callError.ErrorDetails = err.GetDetails()
	return
}

// Get the string representation of the CallError
func (this *CallError) String() string {
	return fmt.Sprintf(
		`[%d, "%s", "%s", "%s", %s]`,
		this.callCode,
		this.UniqueId,
		this.ErrorCode,
		this.ErrorDescription,
		this.ErrorDetails)
}
