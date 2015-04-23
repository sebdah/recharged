package types_test

import (
	"testing"
	"time"

	"github.com/sebdah/recharged/central-system/types"
	"github.com/stretchr/testify/assert"
)

// Test JSON marshalling in UTC
func TestMarshalJSONBasicUTC(t *testing.T) {
	jsonTime := new(types.JSONTime)
	parsedTime, err := time.Parse(time.RFC3339, "2015-12-31T20:00:00Z")
	jsonTime.Time = parsedTime
	assert.Nil(t, err)

	json, err := jsonTime.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, "\"2015-12-31T20:00:00Z\"", string(json))
}

// Test JSON marshalling in CEST
func TestMarshalJSONBasicCEST(t *testing.T) {
	jsonTime := new(types.JSONTime)
	parsedTime, err := time.Parse(time.RFC3339, "2015-12-31T20:00:00+02:00")
	jsonTime.Time = parsedTime
	assert.Nil(t, err)

	json, err := jsonTime.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, "\"2015-12-31T18:00:00Z\"", string(json))
}

// Get the string representation in UTC
func TestStringBasicUTC(t *testing.T) {
	jsonTime := new(types.JSONTime)
	parsed, err := time.Parse(time.RFC3339, "2015-12-31T20:00:00Z")
	jsonTime.Time = parsed
	assert.Nil(t, err)

	assert.Equal(t, "2015-12-31T20:00:00Z", jsonTime.String())
}

// Get the string representation in CEST
func TestStringBasicCEST(t *testing.T) {
	jsonTime := new(types.JSONTime)
	parsed, err := time.Parse(time.RFC3339, "2015-12-31T20:00:00+02:00")
	jsonTime.Time = parsed
	assert.Nil(t, err)

	assert.Equal(t, "2015-12-31T18:00:00Z", jsonTime.String())
}
