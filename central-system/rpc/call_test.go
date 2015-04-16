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
