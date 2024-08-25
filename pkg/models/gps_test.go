package models

import (
	"testing"
	"time"
)

// Usando table tests mesmo em um projeto pequeno pois eles são muito mais fáceis de ler
func TestNewGPSData(t *testing.T) {
	d := &DeviceData{
		MAC:       "AA:BB:CC:DD:EE:FF",
		Timestamp: time.Now(),
	}

	tests := []struct {
		name        string
		deviceData  *DeviceData
		latitude    float64
		longitude   float64
		expectError bool
	}{
		{
			name:        "Valid Device data and Coordinates",
			deviceData:  d,
			latitude:    37.7749,
			longitude:   -122.4194,
			expectError: false,
		},
		{
			name:        "Nil Device data",
			deviceData:  nil,
			latitude:    37.7749,
			longitude:   -122.4194,
			expectError: true,
		},
		{
			name:        "Invalid Latitude",
			deviceData:  d,
			latitude:    100.0, // latitude deve ser entre -90 e 90
			longitude:   -122.4194,
			expectError: true,
		},
		{
			name:        "Invalid Longitude",
			deviceData:  d,
			latitude:    37.7749,
			longitude:   -200.0, // longitude deve ser entre -180 e 180
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gpsData, err := NewGPS(tt.deviceData, tt.latitude, tt.longitude)

			if tt.expectError {
				if err == nil {
					t.Error("Expected an error but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Did not expect an error but got %v", err)
				}
				if gpsData.DeviceData != tt.deviceData {
					t.Errorf("Expected Device %v, got %v", tt.deviceData, gpsData.DeviceData)
				}
				if gpsData.Latitude != tt.latitude {
					t.Errorf("Expected Latitude %f, got %f", tt.latitude, gpsData.Latitude)
				}
				if gpsData.Longitude != tt.longitude {
					t.Errorf("Expected Longitude %f, got %f", tt.longitude, gpsData.Longitude)
				}
			}
		})
	}
}
