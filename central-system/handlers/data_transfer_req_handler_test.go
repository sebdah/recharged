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
	baseUrl = fmt.Sprintf("%s/ocpp/v2.0j/dataTransfer", server.URL)
}

func TestDataTransferBasic(t *testing.T) {
	// Build the DataTransfer.req
	req := new(messages.DataTransferReq)
	req.VendorId = "se.agrea.vendor"

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
	conf := new(messages.DataTransferConf)
	decoder := json.NewDecoder(res.Body)
	_ = decoder.Decode(&conf)

	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, "UnknownVendorId", conf.Status)

	return
}
