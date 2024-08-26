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

func TestCreateGyroscope(t *testing.T) {
	router := setupRouter("/gyroscope")
	now := time.Now()

	tests := []struct {
		name           string
		requestBody    interface{}
		expectedStatus int
		expectedBody   httpcore.ApiError
	}{
		{
			name: "Valid Gyroscope data",
			requestBody: func() interface{} {
				x := float64(37.7749)
				y := float64(-122.4194)
				z := float64(300.1)
				gyro, _ := models.NewGyroscope(
					&models.DeviceData{
						MAC:       "AA:BB:CC:DD:EE:FF",
						Timestamp: now,
					},
					&x,
					&y,
					&z,
				)
				return gyro
			}(),
			expectedStatus: http.StatusCreated,
			expectedBody:   httpcore.ApiError{},
		},
		{
			name: "Missing DeviceData",
			requestBody: func() interface{} {
				x := float64(37.7749)
				y := float64(-122.4194)
				z := float64(300.1)
				gyro := models.Gyroscope{
					X: &x,
					Y: &y,
					Z: &z,
				}
				return gyro
			}(),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   httpcore.ErrBadRequest.With(errors.New("device cannot be nil")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestBody, _ := json.Marshal(tt.requestBody)
			req, _ := http.NewRequest(http.MethodPost, "/telemetry/gyroscope", bytes.NewBuffer(requestBody))
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
