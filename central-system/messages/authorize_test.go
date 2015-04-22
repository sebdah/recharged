package messages

import (
	"testing"

	"github.com/sebdah/recharged/central-system/models"
	"github.com/stretchr/testify/assert"
)

func init() {
	idTag := models.NewIdTag()
	models.DropCollection(idTag)
	models.EnsureIndexes(idTag)
}

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

// WIP - Test processing of missing IdTag
//func TestAuthorizeReqProcessMissingIdTag(t *testing.T) {
//// Create the request
//req := new(AuthorizeReq)
//err := req.ParseJson(`{"idTag": { "id": "11234fsd", "idType": "ISO14443" } }`)
//assert.Nil(t, err)

//_, errorer := req.Process()
//assert.Nil(t, errorer)
//}
