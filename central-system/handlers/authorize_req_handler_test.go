package handlers_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/sebdah/recharged/central-system/messages"
	"github.com/sebdah/recharged/central-system/models"
	"github.com/sebdah/recharged/central-system/routers"
	"github.com/sebdah/recharged/central-system/types"
	"github.com/stretchr/testify/assert"
)

var (
	server       *httptest.Server
	reader       io.Reader
	baseUrl      string
	idTagBaseUrl string
)

// Initializer
func init() {
	server = httptest.NewServer(routers.Router())
	baseUrl = fmt.Sprintf("%s/ocpp/v2.0j/authorize", server.URL)
	idTagBaseUrl = fmt.Sprintf("%s/admin/idTags", server.URL)

	// Prepare the database
	idTag := new(models.IdTag)
	models.DropCollection(idTag)
	models.EnsureIndexes(idTag)
}

// Send authorization request
func authorize(t *testing.T, req messages.AuthorizeReq) (*http.Response, messages.AuthorizeConf) {
	// Build the request
	data, err := json.Marshal(req)
	assert.Nil(t, err)
	reader := strings.NewReader(string(data))

	// Send the request
	r, err := http.NewRequest("POST", baseUrl, reader)
	assert.Nil(t, err)
	res, err := http.DefaultClient.Do(r)
	assert.Nil(t, err)

	// Parse the response
	conf := new(messages.AuthorizeConf)
	decoder := json.NewDecoder(res.Body)
	_ = decoder.Decode(&conf)

	return res, *conf
}

// Helper - Create idTag
func createIdTag(t *testing.T, body string) (res *http.Response) {
	reader := strings.NewReader(body)
	r, err := http.NewRequest("POST", idTagBaseUrl, reader)
	assert.Nil(t, err)
	res, err = http.DefaultClient.Do(r)
	assert.Nil(t, err)
	return
}

// Helper - Delete idTag
func deleteIdTag(t *testing.T, idTag string) (res *http.Response) {
	r, err := http.NewRequest("DELETE", idTagBaseUrl+"/"+idTag, nil)
	assert.Nil(t, err)
	res, err = http.DefaultClient.Do(r)
	assert.Nil(t, err)
	return res
}

// Basic test of authorization
func TestAuthorizeBasic(t *testing.T) {
	// Create IdTag
	createIdTag(t, `{ "idTag": "test" }`)

	// Build the Authorize.req
	req := new(messages.AuthorizeReq)
	idToken := new(types.IdToken)
	idToken.Id = "test"
	req.IdTag = *idToken

	// Send the Authorize.req
	res, conf := authorize(t, *req)
	assert.Equal(t, 200, res.StatusCode)
	assert.Nil(t, conf.IdTagInfo.GroupTagId)
	assert.Equal(t, "en", conf.IdTagInfo.Language)
	assert.Equal(t, "Accepted", conf.IdTagInfo.Status)

	// Delete IdTag
	deleteIdTag(t, "test")
}

// Full test of authorization
func TestAuthorizeFull(t *testing.T) {
	// Create IdTag
	createIdTag(t, `{ "idTag": "test", "idType": "ISO14443" }`)

	// Build the Authorize.req
	req := new(messages.AuthorizeReq)
	idToken := new(types.IdToken)
	idToken.Id = "test"
	req.IdTag = *idToken

	// Send the Authorize.req
	res, _ := authorize(t, *req)
	assert.Equal(t, 200, res.StatusCode)

	// Delete IdTag
	deleteIdTag(t, "test")
}

// Test with unknown IdTag
func TestAuthorizeUnknownIdTag(t *testing.T) {
	// Build the Authorize.req
	req := new(messages.AuthorizeReq)
	idToken := new(types.IdToken)
	idToken.Id = "test"
	req.IdTag = *idToken

	// Send the Authorize.req
	res, _ := authorize(t, *req)
	assert.Equal(t, 404, res.StatusCode)
}