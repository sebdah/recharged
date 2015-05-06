package rpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCall(t *testing.T) {
	call := NewCall()
	assert.Equal(t, "", call.UniqueId)
	assert.Equal(t, "", call.Action)
	assert.Equal(t, "{}", call.Payload)
}

// Test basic population
func TestCallPopulate(t *testing.T) {
	call := NewCall()
	callError := call.Populate(`[2, "1234", "Authorize", {}]`)
	assert.Nil(t, callError)
	assert.Equal(t, "1234", call.UniqueId)
	assert.Equal(t, "Authorize", call.Action)
	assert.Equal(t, "{}", call.Payload)
}

// Test test population with null as payload
func TestCallPopulateWithNullPayload(t *testing.T) {
	call := NewCall()
	callError := call.Populate(`[2, "1234", "Authorize", null]`)
	assert.Nil(t, callError)
	assert.Equal(t, "1234", call.UniqueId)
	assert.Equal(t, "Authorize", call.Action)
	assert.Equal(t, "{}", call.Payload)
}

// Test population of malformatted Call request
func TestCallPopulateInvalidCall(t *testing.T) {
	call := NewCall()
	callError := call.Populate(`dfsd`)
	assert.Equal(t, 4, callError.callCode)
	assert.Equal(t, "", callError.UniqueId)
	assert.Equal(t, "GenericError", callError.ErrorCode)
	assert.Equal(t, "Any other error not covered by the previous ones", callError.ErrorDescription)
	assert.Equal(t, "{\"message\": \"Malformatted message\"}", callError.ErrorDetails)
}

// Test string representation
func TestCallString(t *testing.T) {
	call := NewCall()
	callError := call.Populate(`[2, "1234", "Authorize", {}]`)
	assert.Nil(t, callError)
	assert.Equal(t, `[2, "1234", "Authorize", {}]`, call.String())
}
