package handlers_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/sebdah/recharged/central-system/models"
	"github.com/sebdah/recharged/central-system/routers"
	"github.com/stretchr/testify/assert"
)

var (
	server  *httptest.Server
	reader  io.Reader
	baseUrl string
)

// Initializer
func init() {
	server = httptest.NewServer(routers.Router())
	baseUrl = fmt.Sprintf("%s/admin/idTags", server.URL)
}

// Test creation of IdTag
func TestCreateIdTagSimple(t *testing.T) {
	// Create the tag
	body := `{"idTag": "test"}`
	reader := strings.NewReader(body)

	r, err := http.NewRequest("POST", baseUrl, reader)
	assert.Nil(t, err)

	res, err := http.DefaultClient.Do(r)
	assert.Nil(t, err)
	assert.Equal(t, 201, res.StatusCode)

	// Delete it again
	r, err = http.NewRequest("DELETE", baseUrl+"/test", nil)
	assert.Nil(t, err)
	res, err = http.DefaultClient.Do(r)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)
}

// Test creation of IdTag, full example
func TestCreateIdTagFull(t *testing.T) {
	// Create the tag
	body := `
    {
        "idTag": "test",
        "idType": "ISO14443",
        "language": "en",
        "expiryDate": "2015-12-31T20:00:00Z",
        "groupIdTag": "testGroup"
    }`
	reader := strings.NewReader(body)

	r, err := http.NewRequest("POST", baseUrl, reader)
	assert.Nil(t, err)

	res, err := http.DefaultClient.Do(r)
	assert.Nil(t, err)
	assert.Equal(t, 201, res.StatusCode)

	// Fetch it and match the data
	r, err = http.NewRequest("GET", baseUrl+"/test", nil)
	assert.Nil(t, err)
	res, err = http.DefaultClient.Do(r)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	idTag := new(models.IdTag)
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&idTag)
	assert.Nil(t, err)
	assert.Equal(t, "test", idTag.IdTag)
	assert.Equal(t, "ISO14443", idTag.IdType)
	assert.Equal(t, "en", idTag.Language)
	assert.Equal(t, "2015-12-31T20:00:00Z", idTag.ExpiryDate.String())
	assert.Equal(t, "testGroup", idTag.GroupIdTag)

	// Delete it again
	r, err = http.NewRequest("DELETE", baseUrl+"/test", nil)
	assert.Nil(t, err)
	res, err = http.DefaultClient.Do(r)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)
}

// Test creation of IdTag, without required fields
func TestCreateIdTagMissingRequiredField(t *testing.T) {
	// Create the tag
	body := `{"idType": "ISO14443"}`
	reader := strings.NewReader(body)

	r, err := http.NewRequest("POST", baseUrl, reader)
	assert.Nil(t, err)

	res, err := http.DefaultClient.Do(r)
	assert.Nil(t, err)
	assert.Equal(t, 400, res.StatusCode)
}

// Test creation of IdTag, duplicate
func TestCreateIdTagDuplicate(t *testing.T) {
	// Create the tag
	body := `{"idTag": "test"}`
	reader := strings.NewReader(body)

	r, err := http.NewRequest("POST", baseUrl, reader)
	assert.Nil(t, err)
	res, err := http.DefaultClient.Do(r)
	assert.Nil(t, err)
	assert.Equal(t, 201, res.StatusCode)

	reader = strings.NewReader(body)
	r, err = http.NewRequest("POST", baseUrl, reader)
	assert.Nil(t, err)
	res, err = http.DefaultClient.Do(r)
	assert.Nil(t, err)
	assert.Equal(t, 409, res.StatusCode)

	// Delete it again
	r, err = http.NewRequest("DELETE", baseUrl+"/test", nil)
	assert.Nil(t, err)
	res, err = http.DefaultClient.Do(r)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)
}

// Test get of IdTag, full example
func TestGetIdTagFull(t *testing.T) {
	// Create the tag
	body := `
    {
        "idTag": "test",
        "idType": "ISO14443",
        "language": "en",
        "expiryDate": "2015-12-31T20:00:00Z",
        "groupIdTag": "testGroup"
    }`
	reader := strings.NewReader(body)

	r, err := http.NewRequest("POST", baseUrl, reader)
	assert.Nil(t, err)

	res, err := http.DefaultClient.Do(r)
	assert.Nil(t, err)
	assert.Equal(t, 201, res.StatusCode)

	// Fetch it and match the data
	r, err = http.NewRequest("GET", baseUrl+"/test", nil)
	assert.Nil(t, err)
	res, err = http.DefaultClient.Do(r)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	idTag := new(models.IdTag)
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&idTag)
	assert.Nil(t, err)
	assert.Equal(t, "test", idTag.IdTag)
	assert.Equal(t, "ISO14443", idTag.IdType)
	assert.Equal(t, "en", idTag.Language)
	assert.Equal(t, "2015-12-31T20:00:00Z", idTag.ExpiryDate.String())
	assert.Equal(t, "testGroup", idTag.GroupIdTag)

	// Delete it again
	r, err = http.NewRequest("DELETE", baseUrl+"/test", nil)
	assert.Nil(t, err)
	res, err = http.DefaultClient.Do(r)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)
}

// Test fetching IdTag that does not exist
func TestGetIdTagMissing(t *testing.T) {
	r, err := http.NewRequest("GET", baseUrl+"/test", nil)
	assert.Nil(t, err)
	res, err := http.DefaultClient.Do(r)
	assert.Nil(t, err)
	assert.Equal(t, 404, res.StatusCode)
}

// Test deletion of IdTag
func TestDeleteIdTagSimple(t *testing.T) {
	// Create the tag
	body := `{"idTag": "test"}`
	reader := strings.NewReader(body)

	r, err := http.NewRequest("POST", baseUrl, reader)
	assert.Nil(t, err)

	res, err := http.DefaultClient.Do(r)
	assert.Nil(t, err)
	assert.Equal(t, 201, res.StatusCode)

	// Delete it again
	r, err = http.NewRequest("DELETE", baseUrl+"/test", nil)
	assert.Nil(t, err)
	res, err = http.DefaultClient.Do(r)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)
}

// Test deletion of IdTag that does not exist
func TestDeleteIdTagNotExist(t *testing.T) {
	// Delete it again
	r, err := http.NewRequest("DELETE", baseUrl+"/test", nil)
	assert.Nil(t, err)
	res, err := http.DefaultClient.Do(r)
	assert.Nil(t, err)
	assert.Equal(t, 404, res.StatusCode)
}
