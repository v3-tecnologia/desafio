package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"v3/pkg/httpcore"
	"v3/pkg/models"
)

func TestCreateGPS(t *testing.T) {
	router := setupRouter("/gps")
	now := time.Now()

	tests := []struct {
		name           string
		requestBody    models.GPS
		expectedStatus int
		expectedBody   httpcore.ApiError
	}{
		{
			name: "Valid GPS data",
			requestBody: models.GPS{
				Latitude:  37.7749,
				Longitude: -122.4194,
				DeviceData: &models.DeviceData{
					MAC:       "AA:BB:CC:DD:EE:FF",
					Timestamp: now,
				},
			},
			expectedStatus: http.StatusCreated,
			expectedBody:   httpcore.ApiError{},
		},
		{
			name: "Missing DeviceData",
			requestBody: models.GPS{
				Latitude:  37.7749,
				Longitude: -122.4194,
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   httpcore.ErrBadRequest.With(errors.New("device cannot be nil")),
		},
		{
			name: "Invalid latitude",
			requestBody: models.GPS{
				Latitude:  -95.0,
				Longitude: -122.4194,
				DeviceData: &models.DeviceData{
					MAC:       "AA:BB:CC:DD:EE:FF",
					Timestamp: now,
				},
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   httpcore.ErrBadRequest.With(errors.New("latitude must be between -90 and 90")),
		},
		{
			name: "Invalid longitude",
			requestBody: models.GPS{
				Latitude:  37.7749,
				Longitude: -200.0,
				DeviceData: &models.DeviceData{
					MAC:       "AA:BB:CC:DD:EE:FF",
					Timestamp: now,
				},
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   httpcore.ErrBadRequest.With(errors.New("longitude must be between -180 and 180")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestBody, _ := json.Marshal(tt.requestBody)
			req, _ := http.NewRequest(http.MethodPost, "/telemetry/gps", bytes.NewBuffer(requestBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if tt.expectedStatus != http.StatusCreated {
				var responseBody httpcore.ApiError
				err := json.NewDecoder(w.Body).Decode(&responseBody)
				if err != nil {
					t.Fatalf("failed to decode response body: %v", err)
				}

				t.Logf("Actual Response Body: %+v", responseBody)

				if responseBody != tt.expectedBody {
					t.Errorf("expected body %+v, got %+v", tt.expectedBody, responseBody)
				}
			}
		})
	}
}
