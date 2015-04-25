package handlers

import (
	"fmt"
	"log"

	"github.com/sebdah/recharged/shared/messages"
	"github.com/sebdah/recharged/shared/processor"
	"github.com/sebdah/recharged/shared/rpc"
)

// Handle CALL requests
func CallHandler(msg string, proc processor.Processor) (callResult *rpc.CallResult, callError *rpc.CallError) {
	var err error

	// Populate and validate the fields
	call := rpc.NewCall()
	callError = call.Populate(msg)
	if callError != nil {
		return
	}

	// Find the requested message matching the action
	if call.Action == "Authorize" {
		req := messages.NewAuthorizeReq()
		err = req.ParseJson(call.Payload)
		if err != nil {
			log.Println("Error parsing JSON: %s", err.Error())
			callError = rpc.NewCallError(call.UniqueId, rpc.NewFormationViolation())
			return
		}
		conf, err := proc.ProcessAuthorize(req)
		if err != nil {
			log.Println("Error processing request: %s", err.Error())
			genericError := rpc.NewGenericError()
			genericError.SetDetails(fmt.Sprintf(`{"message": "%s"}`, err.Error()))
			callError = rpc.NewCallError(call.UniqueId, genericError)
			return
		}

		callResult = rpc.NewCallResult(call.UniqueId)
		callResult.SetPayload(conf.String())
		return

	} else if call.Action == "BootNotification" {
		req := messages.NewBootNotificationReq()
		err = req.ParseJson(call.Payload)
		if err != nil {
			log.Println("Error parsing JSON: %s", err.Error())
			callError = rpc.NewCallError(call.UniqueId, rpc.NewFormationViolation())
			return
		}
		conf, err := proc.ProcessBootNotification(req)
		if err != nil {
			log.Println("Error processing request: %s", err.Error())
			genericError := rpc.NewGenericError()
			genericError.SetDetails(fmt.Sprintf(`{"message": "%s"}`, err.Error()))
			callError = rpc.NewCallError(call.UniqueId, genericError)
			return
		}

		callResult = rpc.NewCallResult(call.UniqueId)
		callResult.SetPayload(conf.String())
		return

	} else if call.Action == "DataTransfer" {
		req := messages.NewDataTransferReq()
		err = req.ParseJson(call.Payload)
		if err != nil {
			log.Println("Error parsing JSON: %s", err.Error())
			callError = rpc.NewCallError(call.UniqueId, rpc.NewFormationViolation())
			return
		}
		conf, err := proc.ProcessDataTransfer(req)
		if err != nil {
			log.Println("Error processing request: %s", err.Error())
			genericError := rpc.NewGenericError()
			genericError.SetDetails(fmt.Sprintf(`{"message": "%s"}`, err.Error()))
			callError = rpc.NewCallError(call.UniqueId, genericError)
			return
		}

		callResult = rpc.NewCallResult(call.UniqueId)
		callResult.SetPayload(conf.String())
		return

	} else {
		callError = rpc.NewCallError(call.UniqueId, rpc.NewNotImplementedError())
		return
	}

	return
}
