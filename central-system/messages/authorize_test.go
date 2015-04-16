package messages

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test simple JSON parsing
func TestAuthorizeReqParseJson(t *testing.T) {
	req := new(AuthorizeReq)
	err := req.ParseJson(`{"idTag": { "id": "1234", "idType": "ISO14443" } }`)
	assert.Nil(t, err)
	assert.Equal(t, "1234", req.IdTag.Id)
	assert.Equal(t, "ISO14443", req.IdTag.IdType)
}

// Test string representation
func TestAuthorizeReqString(t *testing.T) {
	req := new(AuthorizeReq)
	err := req.ParseJson(`{"idTag": { "id": "1234", "idType": "ISO14443" } }`)
	assert.Nil(t, err)
	assert.Equal(t, `{"idTag":{"id":"1234","idType":"ISO14443"}}`, req.String())
}
