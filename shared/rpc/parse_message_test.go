package rpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test regular CALL
func TestParseMessageCall(t *testing.T) {
	messageType, err := ParseMessage(`[2, "19223201", "BootNotification", {"chargePointVendor": "VendorX", "chargePointModel": "SingleSocketCharger"}]`)
	assert.Nil(t, err)
	assert.Equal(t, 2, messageType)
}

// Test regular CALLRESULT
func TestParseMessageCallResult(t *testing.T) {
	messageType, err := ParseMessage(`[3, "19223201", {"status":"Accepted", "currentTime":"2013-02-01T20:53:32.486Z", "heartbeatInterval":300}]`)
	assert.Nil(t, err)
	assert.Equal(t, 3, messageType)
}

// Test regular CALLRESULT
func TestParseMessageCallError(t *testing.T) {
	messageType, err := ParseMessage(`[4, "123434", "NotImplemented", "Requested Action is not known by receiver", {}]`)
	assert.Nil(t, err)
	assert.Equal(t, 4, messageType)
}

// Test malformatted message
func TestParseMessageMalformatted(t *testing.T) {
	messageType, err := ParseMessage("dsaf")
	assert.Equal(t, 0, messageType)
	assert.Equal(t, "malformatted message", err.Error())
}

// Test unsupported type, too low
func TestParseMessageUnknownTypeLow(t *testing.T) {
	messageType, err := ParseMessage(`[1, "123434", "NotImplemented", "Requested Action is not known by receiver", {}]`)
	assert.Equal(t, 0, messageType)
	assert.Equal(t, "unknown message type", err.Error())
}

// Test unsupported type, too high
func TestParseMessageUnknownHIgh(t *testing.T) {
	messageType, err := ParseMessage(`[5, "123434", "NotImplemented", "Requested Action is not known by receiver", {}]`)
	assert.Equal(t, 0, messageType)
	assert.Equal(t, "unknown message type", err.Error())
}
