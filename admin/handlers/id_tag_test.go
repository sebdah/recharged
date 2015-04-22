package handlers_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/sebdah/recharged/admin/models"
	"github.com/sebdah/recharged/admin/routers"
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

	// Prepare the database
	idTag := new(models.IdTag)
	models.DropCollection(idTag)
	models.EnsureIndexes(idTag)
}

// Helper - Create idTag
func createIdTag(t *testing.T, body string) (res *http.Response) {
	reader := strings.NewReader(body)
	r, err := http.NewRequest("POST", baseUrl, reader)
	assert.Nil(t, err)
	res, err = http.DefaultClient.Do(r)
	assert.Nil(t, err)
	return
}

// Helper - Delete idTag
func deleteIdTag(t *testing.T, idTag string) (res *http.Response) {
	r, err := http.NewRequest("DELETE", baseUrl+"/"+idTag, nil)
	assert.Nil(t, err)
	res, err = http.DefaultClient.Do(r)
	assert.Nil(t, err)
	return res
}

// Test listing of IdTags
func TestListIdTagSimple(t *testing.T) {
	// Create the tag
	res := createIdTag(t, `{"idTag": "test1"}`)
	assert.Equal(t, 201, res.StatusCode)

	// Create the tag
	res = createIdTag(t, `{"idTag": "test2"}`)
	assert.Equal(t, 201, res.StatusCode)

	// List all IdTags
	r, err := http.NewRequest("GET", baseUrl, reader)
	assert.Nil(t, err)
	res, err = http.DefaultClient.Do(r)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	var idTags []models.IdTag
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&idTags)
	assert.Nil(t, err)
	assert.Len(t, idTags, 2)

	// Delete it
	res = deleteIdTag(t, "test1")
	assert.Equal(t, 200, res.StatusCode)
	res = deleteIdTag(t, "test2")
	assert.Equal(t, 200, res.StatusCode)
}

// Test creation of IdTag
func TestCreateIdTagSimple(t *testing.T) {
	// Create the tag
	res := createIdTag(t, `{"idTag": "test"}`)
	assert.Equal(t, 201, res.StatusCode)

	// Delete it
	res = deleteIdTag(t, "test")
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
	// Create the tag
	res := createIdTag(t, body)
	assert.Equal(t, 201, res.StatusCode)

	// Fetch it and match the data
	r, err := http.NewRequest("GET", baseUrl+"/test", nil)
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

	// Delete it
	res = deleteIdTag(t, "test")
	assert.Equal(t, 200, res.StatusCode)
}

// Test setting of the active flag
func TestCreateIdTagActiveFlag(t *testing.T) {
	// Create the tag
	res := createIdTag(t, `{"idTag": "test1", "active": true}`)
	assert.Equal(t, 201, res.StatusCode)
	res = createIdTag(t, `{"idTag": "test2", "active": false}`)
	assert.Equal(t, 201, res.StatusCode)

	// Fetch it and match the data
	r, err := http.NewRequest("GET", baseUrl+"/test1", nil)
	assert.Nil(t, err)
	res, err = http.DefaultClient.Do(r)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	idTag := new(models.IdTag)
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&idTag)
	assert.Nil(t, err)
	assert.Equal(t, true, idTag.Active)

	// Fetch it and match the data
	r, err = http.NewRequest("GET", baseUrl+"/test2", nil)
	assert.Nil(t, err)
	res, err = http.DefaultClient.Do(r)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	idTag = new(models.IdTag)
	decoder = json.NewDecoder(res.Body)
	err = decoder.Decode(&idTag)
	assert.Nil(t, err)
	assert.Equal(t, false, idTag.Active)

	// Delete it
	res = deleteIdTag(t, "test1")
	assert.Equal(t, 200, res.StatusCode)
	res = deleteIdTag(t, "test2")
	assert.Equal(t, 200, res.StatusCode)
}

// Test creation of IdTag, without required fields
func TestCreateIdTagMissingRequiredField(t *testing.T) {
	// Create the tag
	res := createIdTag(t, `{"idType": "ISO14443"}`)
	assert.Equal(t, 400, res.StatusCode)
}

// Test creation of IdTag, duplicate
func TestCreateIdTagDuplicate(t *testing.T) {
	// Create the tag
	res := createIdTag(t, `{"idTag": "test"}`)
	assert.Equal(t, 201, res.StatusCode)

	// Try again
	res = createIdTag(t, `{"idTag": "test"}`)
	assert.Equal(t, 409, res.StatusCode)

	// Delete it
	res = deleteIdTag(t, "test")
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
	res := createIdTag(t, body)
	assert.Equal(t, 201, res.StatusCode)

	// Fetch it and match the data
	r, err := http.NewRequest("GET", baseUrl+"/test", nil)
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

	// Delete it
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
	res := createIdTag(t, `{"idTag": "test"}`)
	assert.Equal(t, 201, res.StatusCode)

	// Delete it
	r, err := http.NewRequest("DELETE", baseUrl+"/test", nil)
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

// Test updating of IdTag, full example
func TestUpdateIdTag(t *testing.T) {
	// Create the tag
	body := `
    {
        "idTag": "test",
        "idType": "ISO14443",
        "language": "en",
        "expiryDate": "2015-12-31T20:00:00Z",
        "groupIdTag": "testGroup"
    }`
	res := createIdTag(t, body)
	assert.Equal(t, 201, res.StatusCode)

	// Fetch it and match the data
	r, err := http.NewRequest("GET", baseUrl+"/test", nil)
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

	// Update the tag
	body = `
    {
        "idTag": "test",
        "idType": "Local1",
        "language": "sv",
        "expiryDate": "2016-12-31T20:00:00Z",
        "groupIdTag": "testGroup2"
    }`
	reader = strings.NewReader(body)

	r, err = http.NewRequest("PUT", baseUrl+"/test", reader)
	assert.Nil(t, err)
	res, err = http.DefaultClient.Do(r)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	// Fetch it and match the data
	r, err = http.NewRequest("GET", baseUrl+"/test", nil)
	assert.Nil(t, err)
	res, err = http.DefaultClient.Do(r)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	idTag = new(models.IdTag)
	decoder = json.NewDecoder(res.Body)
	err = decoder.Decode(&idTag)
	assert.Nil(t, err)
	assert.Equal(t, "test", idTag.IdTag)
	assert.Equal(t, "Local1", idTag.IdType)
	assert.Equal(t, "sv", idTag.Language)
	assert.Equal(t, "2016-12-31T20:00:00Z", idTag.ExpiryDate.String())
	assert.Equal(t, "testGroup2", idTag.GroupIdTag)

	// Delete it again
	r, err = http.NewRequest("DELETE", baseUrl+"/test", nil)
	assert.Nil(t, err)
	res, err = http.DefaultClient.Do(r)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)
}

// Test update of IdTag that does not exist
func TestUpdateIdTagNotExist(t *testing.T) {
	// Delete it again
	r, err := http.NewRequest("PUT", baseUrl+"/test", nil)
	assert.Nil(t, err)
	res, err := http.DefaultClient.Do(r)
	assert.Nil(t, err)
	assert.Equal(t, 404, res.StatusCode)
}
