package handlers

import (
	"log"

	"github.com/sebdah/recharged/shared/processors"
	"github.com/sebdah/recharged/shared/rpc"
)

// Handle CALLRESULT messages
func CallResultHandler(msg string, confProcessor processors.ConfProcessor) {
	callResult := new(rpc.CallResult)
	err := callResult.Populate(msg)
	if err != nil {
		log.Printf("Error parsing CALLRESULT: %s\n", err)
		return
	}

	log.Printf("CALLRESULT: %s\n", callResult.String())
}
