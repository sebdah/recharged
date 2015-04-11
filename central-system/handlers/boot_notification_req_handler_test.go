package handlers_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/sebdah/recharged/central-system/messages"
	"github.com/sebdah/recharged/central-system/models"
	"github.com/sebdah/recharged/central-system/routers"
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
	baseUrl = fmt.Sprintf("%s/ocpp/v2.0j/bootNotification", server.URL)

	// Prepare the database
	chargePoint := new(models.ChargePoint)
	models.DropCollection(chargePoint)
	models.EnsureIndexes(chargePoint)
}

func sendBootNotification(t *testing.T, req *messages.BootNotificationReq) (res *http.Response, conf *messages.BootNotificationConf) {
	data, err := json.Marshal(req)
	assert.Nil(t, err)
	reader := strings.NewReader(string(data))

	// Send the request
	r, err := http.NewRequest("POST", baseUrl, reader)
	assert.Nil(t, err)
	res, err = http.DefaultClient.Do(r)
	assert.Nil(t, err)

	// Parse the response
	conf = new(messages.BootNotificationConf)
	decoder := json.NewDecoder(res.Body)
	_ = decoder.Decode(&conf)

	return
}

// Basic test of the functionality
func TestBootNotificationBasic(t *testing.T) {
	// Build the request
	req := new(messages.BootNotificationReq)
	req.ChargePointVendor = "MyVendor"
	req.ChargePointModel = "MyModel"
	req.ChargePointSerialNumber = "12345678"
	req.Imsi = "123412341234"

	// Send the request
	res, conf := sendBootNotification(t, req)
	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, 10, conf.HeartbeatInterval)
	assert.True(t, time.Now().UTC().After(conf.CurrentTime.Time))
	assert.Equal(t, "Accepted", conf.Status)
}
