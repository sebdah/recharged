package rpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test constructor
func TestNewCallResult(t *testing.T) {
	callResult := NewCallResult("1234")
	assert.Equal(t, "1234", callResult.UniqueId)
	assert.Equal(t, "{}", callResult.Payload)
}

// Test setting the payload
func TestCallResultUpdatePayload(t *testing.T) {
	callResult := NewCallResult("1234")
	assert.Equal(t, "1234", callResult.UniqueId)
	assert.Equal(t, "{}", callResult.Payload)
	callResult.SetPayload(`{"test": "string"}`)
	assert.Equal(t, `{"test": "string"}`, callResult.Payload)
}

// Test the string representation
func TestCallResultString(t *testing.T) {
	callResult := NewCallResult("1234")
	callResult.SetPayload(`{"test": "string"}`)
	assert.Equal(t, `[3, "1234", {"test": "string"}]`, callResult.String())
}
