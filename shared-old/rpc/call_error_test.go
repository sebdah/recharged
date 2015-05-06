package rpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test constructor
func TestNewCallError(t *testing.T) {
	callError := NewCallError("1234", NewGenericError())
	assert.Equal(t, 4, callError.callCode)
	assert.Equal(t, "1234", callError.UniqueId)
	assert.Equal(t, "GenericError", callError.ErrorCode)
	assert.Equal(t, "Any other error not covered by the previous ones", callError.ErrorDescription)
	assert.Equal(t, "{}", callError.ErrorDetails)
}

// Test string representation
func TestCallErrorString(t *testing.T) {
	callError := NewCallError("1234", NewGenericError())
	res := callError.String()
	assert.Equal(t, `[4, "1234", "GenericError", "Any other error not covered by the previous ones", {}]`, res)
}
