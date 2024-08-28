package handlers_test

import (
	"github/desafio/handlers"
	"github/desafio/models"
	"github/desafio/service/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type HandlersTestSuite struct {
	suite.Suite
	handle  handlers.Handle
	service *mocks.ProcessData
}

func TestHandlersTestSuite(t *testing.T) {
	suite.Run(t, new(HandlersTestSuite))
}

func (ts *HandlersTestSuite) SetupTest() {
	ts.service = &mocks.ProcessData{}
	ts.handle.Service = ts.service
}

func (ts *HandlersTestSuite) TestGyroscopeData() {
	ts.Run("Success: all data corrected", func() {
		data := models.Gyroscope{
			MacAddress: "00:00:00:00:00:00",
			X:          123.1,
			Y:          213.2,
			Z:          32.3,
			Timestamp:  1724855500}
		ts.service.On("ProcessGyroscopeData", data).Return(nil)

		body := strings.NewReader(`{
			"macAddress" : "00:00:00:00:00:00",
			"x" : 123.1,
			"y" : 213.2,
			"z" : 32.3,
			"timeStamp" :  1724855500
			}`)

		req, err := http.NewRequest("POST", "/telemetry/gyroscope", body)
		ts.Nil(err)

		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(ts.handle.GyroscopeData)

		handler.ServeHTTP(resp, req)

		ts.Equal(http.StatusOK, resp.Code)
	})

	ts.Run("Fail: Empty body", func() {
		req, err := http.NewRequest("POST", "/telemetry/gyroscope", nil)
		ts.Nil(err)

		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(ts.handle.GyroscopeData)

		handler.ServeHTTP(resp, req)

		ts.Equal(http.StatusBadRequest, resp.Code)
	})

	ts.Run("Fail: Invalid data", func() {
		body := strings.NewReader(
			`{
		"macAddress" : "00:B0:C1:75L.26",
		"timeStamp" : 1724603773,`)

		req, err := http.NewRequest("POST", "/telemetry/gyroscope", body)
		ts.Nil(err)

		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(ts.handle.GyroscopeData)

		handler.ServeHTTP(resp, req)

		ts.Equal(http.StatusBadRequest, resp.Code)
	})
}

func (ts *HandlersTestSuite) TestGPSData() {
	ts.Run("Success: all data corrected", func() {
		data := models.GPS{
			MacAddress: "00:00:00:00:00:00",
			Latitude:   "-5.088889",
			Longitude:  "-42.801944",
			Timestamp:  1724855500}
		ts.service.On("ProcessGPSData", data).Return(nil)

		body := strings.NewReader(`{
			"macAddress" : "00:00:00:00:00:00",
			"latitude" : "-5.088889",
			"longitude" : "-42.801944",
			"timeStamp" :  1724855500
			}`)

		req, err := http.NewRequest("POST", "/telemetry/gps", body)
		ts.Nil(err)

		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(ts.handle.GPSData)

		handler.ServeHTTP(resp, req)

		ts.Equal(http.StatusOK, resp.Code)
	})

	ts.Run("Fail: Empty body", func() {
		req, err := http.NewRequest("POST", "/telemetry/gps", nil)
		ts.Nil(err)

		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(ts.handle.GPSData)

		handler.ServeHTTP(resp, req)

		ts.Equal(http.StatusBadRequest, resp.Code)
	})

	ts.Run("Fail: Invalid data", func() {
		body := strings.NewReader(
			`{
		"macAddress" : "00:B0:C1:75L.26",
		"timeStamp" : 1724603773,`)

		req, err := http.NewRequest("POST", "/telemetry/gyroscope", body)
		ts.Nil(err)

		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(ts.handle.GPSData)

		handler.ServeHTTP(resp, req)

		ts.Equal(http.StatusBadRequest, resp.Code)
	})
}

func (ts *HandlersTestSuite) TestPhotoData() {
	ts.Run("Success: all data corrected", func() {
		data := models.Photo{
			MacAddress: "00:00:00:00:00:00",
			Photo:      handlers.Photo,
			Timestamp:  1724855500}
		ts.service.On("ProcessPhoto", data).Return(nil)

		body := strings.NewReader(handlers.JsonTest)

		req, err := http.NewRequest("POST", "/telemetry/photo", body)
		ts.Nil(err)

		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(ts.handle.PhotoData)

		handler.ServeHTTP(resp, req)

		ts.Equal(http.StatusOK, resp.Code)
	})

	ts.Run("Fail: Empty body", func() {
		req, err := http.NewRequest("POST", "/telemetry/photo", nil)
		ts.Nil(err)

		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(ts.handle.GPSData)

		handler.ServeHTTP(resp, req)

		ts.Equal(http.StatusBadRequest, resp.Code)
	})

	ts.Run("Fail: Invalid data", func() {
		body := strings.NewReader(
			`{
		"macAddress" : "00:B0:C1:75L.26",
		"timeStamp" : 1724603773,`)

		req, err := http.NewRequest("POST", "/telemetry/photo", body)
		ts.Nil(err)

		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(ts.handle.GPSData)

		handler.ServeHTTP(resp, req)

		ts.Equal(http.StatusBadRequest, resp.Code)
	})
}