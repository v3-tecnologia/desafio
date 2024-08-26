package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheck)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestPhotoHandler(t *testing.T) {
	reqBody := strings.NewReader(jsonEx)

	reqInvalidBody := strings.NewReader(fmt.Sprintf(
		`{
	"macAddr" : "00-B0-D0-6l2.26",
	"timeStamp" : 1724603773,`))

	reqOK, err := http.NewRequest("POST", "/telemetry/photo", reqBody)
	if err != nil {
		t.Fatal(err)
	}

	reqNoBody, err := http.NewRequest("POST", "/telemetry/photo", nil)
	if err != nil {
		t.Fatal(err)
	}

	reqInvalidData, err := http.NewRequest("POST", "/telemetry/photo", reqInvalidBody)
	if err != nil {
		t.Fatal(err)
	}

	recorderOK := httptest.NewRecorder()
	recorderNoBody := httptest.NewRecorder()
	recorderInvalidData := httptest.NewRecorder()
	handler := http.HandlerFunc(PhotoHandler)

	handler.ServeHTTP(recorderOK, reqOK)
	if status := recorderOK.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	handler.ServeHTTP(recorderNoBody, reqNoBody)
	if status := recorderNoBody.Code; status != http.StatusBadRequest {
		t.Errorf("handler return wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	handler.ServeHTTP(recorderInvalidData, reqInvalidData)
	if status := recorderInvalidData.Code; status != http.StatusBadRequest {
		t.Errorf("handler return wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

}
func TestGpsHandler(t *testing.T) {
	reqBody := strings.NewReader(fmt.Sprintf(
		`{
	"macAddr" : "00-B0-D0-63-C2-26",
	"latitude" : "-18.909762",
	"longitude" : "-48.232750",
	"timeStamp" : 1724603773
	}`))

	reqInvalidBody := strings.NewReader(fmt.Sprintf(
		`{
	"macAddr" : "00-B0.D0-63-C226",
	"latitude" : "-18.909762",
	"timeStamp" : 1724603773
	}`))

	reqOK, err := http.NewRequest("POST", "/telemetry/gps", reqBody)
	if err != nil {
		t.Fatal(err)
	}

	reqNoBody, err := http.NewRequest("POST", "/telemetry/gps", nil)
	if err != nil {
		t.Fatal(err)
	}

	reqInvalidData, err := http.NewRequest("POST", "/telemetry/gps", reqInvalidBody)
	if err != nil {
		t.Fatal(err)
	}

	recorderOK := httptest.NewRecorder()
	recorderNoBody := httptest.NewRecorder()
	recorderInvalidData := httptest.NewRecorder()
	handler := http.HandlerFunc(GpsHandler)

	handler.ServeHTTP(recorderOK, reqOK)
	if status := recorderOK.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	handler.ServeHTTP(recorderNoBody, reqNoBody)
	if status := recorderNoBody.Code; status != http.StatusBadRequest {
		t.Errorf("handler return wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	handler.ServeHTTP(recorderInvalidData, reqInvalidData)
	if status := recorderInvalidData.Code; status != http.StatusBadRequest {
		t.Errorf("handler return wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

}

func TestGyroscopeHandler(t *testing.T) {
	reqBody := strings.NewReader(fmt.Sprintf(`{
	"macAddr" : "00-B0-D0-63-C2-26",
	"x" : 111.2,
	"y" : 222.3,
	"z" : 333.4,
	"timeStamp" : 1724603773
	}`))

	reqInvalidBody := strings.NewReader(fmt.Sprintf(`{
	"macAddr" : "00-B0.D0-63-26",
	"x" : 111.2,
	"z" : 333.4,
	"timeStamp" : 1724603773
	}`))

	reqOK, err := http.NewRequest("POST", "/telemetry/gyroscope", reqBody)
	if err != nil {
		t.Fatal(err)
	}

	reqNoBody, err := http.NewRequest("POST", "/telemetry/gyroscope", nil)
	if err != nil {
		t.Fatal(err)
	}

	reqInvalidData, err := http.NewRequest("POST", "/telemetry/gyroscope", reqInvalidBody)
	if err != nil {
		t.Fatal(err)
	}

	recorderOK := httptest.NewRecorder()
	recorderNoBody := httptest.NewRecorder()
	recorderInvalidData := httptest.NewRecorder()
	handler := http.HandlerFunc(GyroscopeHandler)

	handler.ServeHTTP(recorderOK, reqOK)
	if status := recorderOK.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	handler.ServeHTTP(recorderNoBody, reqNoBody)
	if status := recorderNoBody.Code; status != http.StatusBadRequest {
		t.Errorf("handler return wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	handler.ServeHTTP(recorderInvalidData, reqInvalidData)
	if status := recorderInvalidData.Code; status != http.StatusBadRequest {
		t.Errorf("handler return wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}
