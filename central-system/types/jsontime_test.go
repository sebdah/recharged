package types_test

import (
	"testing"
	"time"

	"github.com/sebdah/recharged/central-system/types"
	"github.com/stretchr/testify/assert"
)

func TestMarshalJSONBasicUTC(t *testing.T) {
	jsonTime := new(types.JSONTime)
	time, err := time.Parse(time.RFC3339, "2015-12-31T20:00:00Z")
	jsonTime.Time = time
	assert.Nil(t, err)
	json, err := jsonTime.Time.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, "\"2015-12-31T20:00:00Z\"", string(json))
}

func TestMarshalJSONBasicCEST(t *testing.T) {
	jsonTime := new(types.JSONTime)
	parsedTime, err := time.Parse(time.RFC3339, "2015-12-31T20:00:00+02:00")
	jsonTime.Time = parsedTime
	assert.Nil(t, err)
	json, err := jsonTime.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, "\"2015-12-31T18:00:00Z\"", string(json))
}
