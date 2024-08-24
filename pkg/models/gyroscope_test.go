package models

import (
	"testing"
	"time"
)

func TestNewGyroscopeData(t *testing.T) {
	d := &DeviceData{
		MAC:       "AA:BB:CC:DD:EE:FF",
		Timestamp: time.Now(),
	}

	tests := []struct {
		name        string
		device      *DeviceData
		x           float64
		y           float64
		z           float64
		expectError bool
	}{
		{
			name:        "Valid Device and Gyroscope Data",
			device:      d,
			x:           1.0,
			y:           0.5,
			z:           -0.5,
			expectError: false,
		},
		{
			name:        "Nil Device data",
			device:      nil,
			x:           1.0,
			y:           0.5,
			z:           -0.5,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gyroData, err := NewGyroscope(tt.device, tt.x, tt.y, tt.z)

			if tt.expectError {
				if err == nil {
					t.Error("Expected an error but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Did not expect an error but got %v", err)
				}
				if gyroData.DeviceData != tt.device {
					t.Errorf("Expected Device %v, got %v", tt.device, gyroData.DeviceData)
				}
				if gyroData.X != tt.x {
					t.Errorf("Expected X %f, got %f", tt.x, gyroData.X)
				}
				if gyroData.Y != tt.y {
					t.Errorf("Expected Y %f, got %f", tt.y, gyroData.Y)
				}
				if gyroData.Z != tt.z {
					t.Errorf("Expected Z %f, got %f", tt.z, gyroData.Z)
				}
			}
		})
	}
}
