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

var chargePointsBaseUrl string

// Initializer
func init() {
	server = httptest.NewServer(routers.Router())
	chargePointsBaseUrl = fmt.Sprintf("%s/chargepoints", server.URL)

	// Prepare the database
	chargePoint := models.NewChargePoint()
	models.DropCollection(chargePoint)
	models.EnsureIndexes(chargePoint)
}

// Helper create ChargePoint
func createChargePoint(t *testing.T, body string) (res *http.Response, chargePoint *models.ChargePoint) {
	reader := strings.NewReader(body)
	fmt.Println(chargePointsBaseUrl)
	r, err := http.NewRequest("POST", chargePointsBaseUrl, reader)
	assert.Nil(t, err)
	res, err = http.DefaultClient.Do(r)
	assert.Nil(t, err)

	// Get the body
	chargePoint = models.NewChargePoint()
	decoder := json.NewDecoder(res.Body)
	_ = decoder.Decode(&chargePoint)

	return
}

// Helper - Delete ChargePoint
func deleteChargePoint(t *testing.T, id string) (res *http.Response) {
	r, err := http.NewRequest("DELETE", chargePointsBaseUrl+"/"+id, nil)
	assert.Nil(t, err)
	res, err = http.DefaultClient.Do(r)
	assert.Nil(t, err)
	return res
}

// Test listing of ChargePoints
func TestListChargePointSimple(t *testing.T) {
	// Create the chargePoint
	res, chargePoint1 := createChargePoint(t, `{"model": "model1", "vendor": "vendor1"}`)
	assert.Equal(t, 201, res.StatusCode)

	// Create the chargePoint
	res, chargePoint2 := createChargePoint(t, `{"model": "model2", "vendor": "vendor2"}`)
	assert.Equal(t, 201, res.StatusCode)

	// List all ChargePoints
	r, err := http.NewRequest("GET", chargePointsBaseUrl, reader)
	assert.Nil(t, err)
	res, err = http.DefaultClient.Do(r)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	var chargePoints []models.ChargePoint
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&chargePoints)
	assert.Nil(t, err)
	assert.Len(t, chargePoints, 2)

	// Delete it
	res = deleteChargePoint(t, chargePoint1.Id.Hex())
	assert.Equal(t, 200, res.StatusCode)
	res = deleteChargePoint(t, chargePoint2.Id.Hex())
	assert.Equal(t, 200, res.StatusCode)
}

// Test creation of ChargePoint
func TestCreateChargePointSimple(t *testing.T) {
	// Create the chargePoint
	res, chargePoint := createChargePoint(t, `{"model": "model1", "vendor": "vendor1"}`)
	assert.Equal(t, 201, res.StatusCode)

	// Delete it
	res = deleteChargePoint(t, chargePoint.Id.Hex())
	assert.Equal(t, 200, res.StatusCode)
}

// Test creation of ChargePoint, full example
func TestCreateChargePointFull(t *testing.T) {
	// Create the chargePoint
	body := `
		{
			"model": "model1",
			"vendor": "vendor1",
			"serialNumber": "1234",
			"imsi": "5678"
		}`
	// Create the chargePoint
	res, chargePoint1 := createChargePoint(t, body)
	assert.Equal(t, 201, res.StatusCode)

	// Fetch it and match the data
	r, err := http.NewRequest("GET", chargePointsBaseUrl+"/"+chargePoint1.Id.Hex(), nil)
	assert.Nil(t, err)
	res, err = http.DefaultClient.Do(r)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	chargePoint2 := new(models.ChargePoint)
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&chargePoint2)
	assert.Nil(t, err)
	assert.Equal(t, chargePoint1.Model, chargePoint2.Model)
	assert.Equal(t, chargePoint1.Vendor, chargePoint2.Vendor)
	assert.Equal(t, chargePoint1.SerialNumber, chargePoint2.SerialNumber)
	assert.Equal(t, chargePoint1.Imsi, chargePoint2.Imsi)

	// Delete it
	res = deleteChargePoint(t, chargePoint1.Id.Hex())
	assert.Equal(t, 200, res.StatusCode)
}

// Test creation of ChargePoint, missing model
func TestCreateChargePointMissingModel(t *testing.T) {
	// Create the chargePoint
	res, _ := createChargePoint(t, `{"vendor": "vendor1"}`)
	assert.Equal(t, 400, res.StatusCode)
}

// Test creation of ChargePoint, missing vendor
func TestCreateChargePointMissingVendor(t *testing.T) {
	// Create the chargePoint
	res, _ := createChargePoint(t, `{"model": "model1"}`)
	assert.Equal(t, 400, res.StatusCode)
}

// Test get of ChargePoint, full example
func TestGetChargePointFull(t *testing.T) {
	// Create the chargePoint
	body := `
		{
			"model": "model1",
			"vendor": "vendor1",
			"serialNumber": "sn1",
			"imsi": "imsi1"
		}`
	res, chargePoint1 := createChargePoint(t, body)
	assert.Equal(t, 201, res.StatusCode)

	// Fetch it and match the data
	r, err := http.NewRequest("GET", chargePointsBaseUrl+"/"+chargePoint1.Id.Hex(), nil)
	assert.Nil(t, err)
	res, err = http.DefaultClient.Do(r)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	chargePoint2 := new(models.ChargePoint)
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&chargePoint2)
	assert.Nil(t, err)
	assert.Equal(t, chargePoint1.Model, chargePoint2.Model)
	assert.Equal(t, chargePoint1.Vendor, chargePoint2.Vendor)
	assert.Equal(t, chargePoint1.SerialNumber, chargePoint2.SerialNumber)
	assert.Equal(t, chargePoint1.Imsi, chargePoint2.Imsi)

	// Delete it
	res = deleteChargePoint(t, chargePoint1.Id.Hex())
	assert.Equal(t, 200, res.StatusCode)
}

// Test fetching ChargePoint that does not exist
func TestGetChargePointMissing(t *testing.T) {
	r, err := http.NewRequest("GET", chargePointsBaseUrl+"/"+"12345", nil)
	assert.Nil(t, err)
	res, err := http.DefaultClient.Do(r)
	assert.Nil(t, err)
	assert.Equal(t, 404, res.StatusCode)
}

// Test deletion of ChargePoint
func TestDeleteChargePointSimple(t *testing.T) {
	// Create the chargePoint
	res, chargePoint := createChargePoint(t, `{"model": "model1", "vendor": "vendor1"}`)
	assert.Equal(t, 201, res.StatusCode)

	// Delete it
	res = deleteChargePoint(t, chargePoint.Id.Hex())
	assert.Equal(t, 200, res.StatusCode)
}

// Test deletion of ChargePoint that does not exist
func TestDeleteChargePointNotExist(t *testing.T) {
	res := deleteChargePoint(t, "1234")
	assert.Equal(t, 404, res.StatusCode)
}

// Test updating of ChargePoint, full example
func TestUpdateChargePoint(t *testing.T) {
	// Create the chargePoint
	body := `
		{
			"model": "model1",
			"vendor": "vendor1",
			"serialNumber": "sn1",
			"imsi": "imsi1"
		}`
	res, chargePoint1 := createChargePoint(t, body)
	assert.Equal(t, 201, res.StatusCode)

	// Fetch it and match the data
	r, err := http.NewRequest("GET", chargePointsBaseUrl+"/"+chargePoint1.Id.Hex(), nil)
	assert.Nil(t, err)
	res, err = http.DefaultClient.Do(r)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	chargePoint2 := new(models.ChargePoint)
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&chargePoint2)
	assert.Nil(t, err)
	assert.Equal(t, chargePoint1.Model, chargePoint2.Model)
	assert.Equal(t, chargePoint1.Vendor, chargePoint2.Vendor)
	assert.Equal(t, chargePoint1.SerialNumber, chargePoint2.SerialNumber)
	assert.Equal(t, chargePoint1.Imsi, chargePoint2.Imsi)

	// Update the chargePoint
	body = `
		{
			"model": "model2",
			"vendor": "vendor2",
			"serialNumber": "sn2",
			"imsi": "imsi2"
		}`
	reader = strings.NewReader(body)

	r, err = http.NewRequest("PUT", chargePointsBaseUrl+"/"+chargePoint1.Id.Hex(), reader)
	assert.Nil(t, err)
	res, err = http.DefaultClient.Do(r)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	// Fetch it and match the data
	r, err = http.NewRequest("GET", chargePointsBaseUrl+"/"+chargePoint1.Id.Hex(), nil)
	assert.Nil(t, err)
	res, err = http.DefaultClient.Do(r)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	chargePoint3 := new(models.ChargePoint)
	decoder = json.NewDecoder(res.Body)
	err = decoder.Decode(&chargePoint3)
	assert.Nil(t, err)
	assert.Equal(t, "model2", chargePoint3.Model)
	assert.Equal(t, "vendor2", chargePoint3.Vendor)
	assert.Equal(t, "sn2", chargePoint3.SerialNumber)
	assert.Equal(t, "imsi2", chargePoint3.Imsi)

	// Delete it again
	r, err = http.NewRequest("DELETE", chargePointsBaseUrl+"/"+chargePoint1.Id.Hex(), nil)
	assert.Nil(t, err)
	res, err = http.DefaultClient.Do(r)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)
}

// Test update of ChargePoint that does not exist
func TestUpdateChargePointNotExist(t *testing.T) {
	// Delete it again
	r, err := http.NewRequest("PUT", chargePointsBaseUrl+"/"+"test", nil)
	assert.Nil(t, err)
	res, err := http.DefaultClient.Do(r)
	assert.Nil(t, err)
	assert.Equal(t, 404, res.StatusCode)
}

// Test validation of model/vendor combination
func TestChargePointValidationOK(t *testing.T) {
	// Create the chargePoint
	body := `
		{
			"model": "model1",
			"vendor": "vendor1"
		}`
	res, chargePoint := createChargePoint(t, body)
	assert.Equal(t, 201, res.StatusCode)

	// See if it exists
	r, err := http.NewRequest("GET", chargePointsBaseUrl+"/validate/vendor1/model1", nil)
	assert.Nil(t, err)
	res, err = http.DefaultClient.Do(r)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	// Delete it
	res = deleteChargePoint(t, chargePoint.Id.Hex())
	assert.Equal(t, 200, res.StatusCode)
}

// Test validation of model/vendor combination
func TestChargePointValidationFail(t *testing.T) {
	// See if it exists
	r, err := http.NewRequest("GET", chargePointsBaseUrl+"/validate/vendor1/model1", nil)
	assert.Nil(t, err)
	res, err := http.DefaultClient.Do(r)
	assert.Nil(t, err)
	assert.Equal(t, 404, res.StatusCode)
}
