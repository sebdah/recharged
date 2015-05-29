package handlers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/sebdah/recharged/admin/models"
	"github.com/sebdah/recharged/admin/routers"
	"github.com/stretchr/testify/assert"
)

var bootNotificationLogBaseUrl string

// Initializer
func init() {
	fmt.Println("Starting BootNotificationLog tests")
	server := httptest.NewServer(routers.Router())
	bootNotificationLogBaseUrl = fmt.Sprintf("%s/bootnotificationlogs", server.URL)

	// Prepare the database
	bootNotificationLog := models.NewBootNotificationLog()
	models.DropCollection(bootNotificationLog)
	models.EnsureIndexes(bootNotificationLog)
}

// Helper create BootNotificationLog
func createBootNotificationLog(t *testing.T, body string) (res *http.Response, bootNotificationLog *models.BootNotificationLog) {
	reader := strings.NewReader(body)
	r, err := http.NewRequest("POST", bootNotificationLogBaseUrl, reader)
	assert.Nil(t, err)
	res, err = http.DefaultClient.Do(r)
	assert.Nil(t, err)

	// Get the body
	bootNotificationLog = models.NewBootNotificationLog()
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&bootNotificationLog)
	assert.Nil(t, err)

	return
}

// Test basic bootNotificationLog creation
func TestCreateBootNotificationLogSimple(t *testing.T) {
	fmt.Println("BootNotificationLog - TestCreateBootNotificationLogSimple")
	input := `{"chargePoint": {"model": "Model X", "vendor": "Vendor Y"}}`
	res, bootNotificationLog := createBootNotificationLog(t, input)
	assert.Equal(t, 201, res.StatusCode)
	assert.NotNil(t, bootNotificationLog.Id)
	assert.Equal(t, "Model X", bootNotificationLog.ChargePoint.Model)
	assert.Equal(t, "Vendor Y", bootNotificationLog.ChargePoint.Vendor)
	assert.Equal(t, "", bootNotificationLog.ChargePoint.SerialNumber)
	assert.Equal(t, "", bootNotificationLog.ChargePoint.Imsi)
	assert.NotNil(t, bootNotificationLog.Ts)
}

// Full test of the creation
func TestCreateBootNotificationLogFull(t *testing.T) {
	fmt.Println("BootNotificationLog - TestCreateBootNotificationLogFull")
	input := `{"chargePoint": {"model": "Model X", "vendor": "Vendor Y", "serialNumber": "1234", "imsi": "12344"}}`
	res, bootNotificationLog := createBootNotificationLog(t, input)
	assert.Equal(t, 201, res.StatusCode)
	assert.NotNil(t, bootNotificationLog.Id)
	assert.Equal(t, "Model X", bootNotificationLog.ChargePoint.Model)
	assert.Equal(t, "Vendor Y", bootNotificationLog.ChargePoint.Vendor)
	assert.Equal(t, "1234", bootNotificationLog.ChargePoint.SerialNumber)
	assert.Equal(t, "12344", bootNotificationLog.ChargePoint.Imsi)
	assert.NotNil(t, bootNotificationLog.Ts)
}

// Test basic bootNotificationLog fetching
func TestGetBootNotificationLog(t *testing.T) {
	// Create a boot notification log
	fmt.Println("BootNotificationLog - TestGetBootNotificationLog")
	input := `{"chargePoint": {"model": "Model A", "vendor": "Vendor B"}}`
	res, bootNotificationLog1 := createBootNotificationLog(t, input)
	assert.Equal(t, 201, res.StatusCode)
	assert.NotNil(t, bootNotificationLog1.Id)

	// Fetch it
	r, err := http.NewRequest("GET", bootNotificationLogBaseUrl+"/"+bootNotificationLog1.Id.Hex(), nil)
	assert.Nil(t, err)
	res, err = http.DefaultClient.Do(r)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	// Match the data
	bootNotificationLog2 := new(models.BootNotificationLog)
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&bootNotificationLog2)
	assert.Equal(t, bootNotificationLog1.ChargePoint.Model, bootNotificationLog2.ChargePoint.Model)
	assert.Equal(t, bootNotificationLog1.ChargePoint.Vendor, bootNotificationLog2.ChargePoint.Vendor)
	assert.Equal(t, bootNotificationLog1.ChargePoint.SerialNumber, bootNotificationLog2.ChargePoint.SerialNumber)
	assert.Equal(t, bootNotificationLog1.ChargePoint.Imsi, bootNotificationLog2.ChargePoint.Imsi)
}
