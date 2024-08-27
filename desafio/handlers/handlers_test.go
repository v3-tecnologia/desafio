package handlers

import (
	"desafio/models"
	"desafio/service/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type HandlersTestSuite struct {
	suite.Suite
	requestHandle RequestHandle
	service       *mocks.IService
}

func TestHandlersTestSuite(t *testing.T) {
	suite.Run(t, new(HandlersTestSuite))
}

func (ts *HandlersTestSuite) SetupTest() {
	ts.service = &mocks.IService{}
	ts.requestHandle.serv = ts.service
}

func (ts *HandlersTestSuite) TestHealthCheckHandler() {
	ts.Run("sucess", func() {
		req, err := http.NewRequest("GET", "/", nil)
		ts.Nil(err)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(ts.requestHandle.HealthCheck)

		handler.ServeHTTP(rr, req)

		ts.Equal(http.StatusOK, rr.Code)
	})
}

func (ts *HandlersTestSuite) TestPhotoHandler() {
	ts.Run("sucess: normal conditions", func() {
		data := models.PhotoRequest{
			Mac:           "00-00-00-00-00-00",
			ImageBase64:   image,
			UnixtimeStamp: 1724603773}
		ts.service.On("ProcessPhotoData", data).Return(nil)

		reqBody := strings.NewReader(jsonEx)

		req, err := http.NewRequest("POST", "/telemetry/photo", reqBody)
		ts.Nil(err)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(ts.requestHandle.PhotoHandler)

		handler.ServeHTTP(rr, req)

		ts.Equal(http.StatusOK, rr.Code)
	})

	ts.Run("fail: empty body", func() {
		req, err := http.NewRequest("POST", "/telemetry/photo", nil)
		ts.Nil(err)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(ts.requestHandle.PhotoHandler)

		handler.ServeHTTP(rr, req)

		ts.Equal(http.StatusBadRequest, rr.Code)
	})

	ts.Run("fail: invalid data", func() {
		reqBody := strings.NewReader(
			`{
		"macAddr" : "00-B0-D0-6l2.26",
		"timeStamp" : 1724603773,`)

		req, err := http.NewRequest("POST", "/telemetry/photo", reqBody)
		ts.Nil(err)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(ts.requestHandle.PhotoHandler)

		handler.ServeHTTP(rr, req)

		ts.Equal(http.StatusBadRequest, rr.Code)
	})
}
func (ts *HandlersTestSuite) TestGpsHandler() {
	ts.Run("sucess: normal conditions", func() {
		data := models.GpsRequest{
			Mac:           "00-00-00-00-00-00",
			Lat:           "-18.909762",
			Lon:           "-48.232750",
			UnixtimeStamp: 1724603773}
		ts.service.On("ProcessGpsData", data).Return(nil)

		reqBody := strings.NewReader(
			`{
		"macAddr" : "00-00-00-00-00-00",
		"latitude" : "-18.909762",
		"longitude" : "-48.232750",
		"timeStamp" : 1724603773
		}`)

		req, err := http.NewRequest("POST", "/telemetry/gps", reqBody)
		ts.Nil(err)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(ts.requestHandle.GpsHandler)

		handler.ServeHTTP(rr, req)

		ts.Equal(http.StatusOK, rr.Code)
	})

	ts.Run("fail: empty body", func() {
		req, err := http.NewRequest("POST", "/telemetry/gps", nil)
		ts.Nil(err)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(ts.requestHandle.GpsHandler)

		handler.ServeHTTP(rr, req)

		ts.Equal(http.StatusBadRequest, rr.Code)
	})

	ts.Run("fail: invalid data", func() {
		reqBody := strings.NewReader(
			`{
		"macAddr" : "00-B0-D0-6l2.26",
		"timeStamp" : 1724603773,`)

		req, err := http.NewRequest("POST", "/telemetry/gps", reqBody)
		ts.Nil(err)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(ts.requestHandle.GpsHandler)

		handler.ServeHTTP(rr, req)

		ts.Equal(http.StatusBadRequest, rr.Code)
	})
}

func (ts *HandlersTestSuite) TestGyroscopeHandler() {
	ts.Run("sucess: normal conditions", func() {
		data := models.GyroscopeRequest{
			Mac:           "00-00-00-00-00-00",
			X:             111.1,
			Y:             222.2,
			Z:             333.3,
			UnixtimeStamp: 1724603773}
		ts.service.On("ProcessGyroscopeData", data).Return(nil)

		reqBody := strings.NewReader(`{
			"macAddr" : "00-00-00-00-00-00",
			"x" : 111.1,
			"y" : 222.2,
			"z" : 333.3,
			"timeStamp" : 1724603773
			}`)

		req, err := http.NewRequest("POST", "/telemetry/gyroscope", reqBody)
		ts.Nil(err)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(ts.requestHandle.GyroscopeHandler)

		handler.ServeHTTP(rr, req)

		ts.Equal(http.StatusOK, rr.Code)
	})

	ts.Run("fail: empty body", func() {
		req, err := http.NewRequest("POST", "/telemetry/gyroscope", nil)
		ts.Nil(err)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(ts.requestHandle.GyroscopeHandler)

		handler.ServeHTTP(rr, req)

		ts.Equal(http.StatusBadRequest, rr.Code)
	})

	ts.Run("fail: invalid data", func() {
		reqBody := strings.NewReader(
			`{
		"macAddr" : "00-B0-D0-6l2.26",
		"timeStamp" : 1724603773,`)

		req, err := http.NewRequest("POST", "/telemetry/gyroscope", reqBody)
		ts.Nil(err)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(ts.requestHandle.GyroscopeHandler)

		handler.ServeHTTP(rr, req)

		ts.Equal(http.StatusBadRequest, rr.Code)
	})
}
