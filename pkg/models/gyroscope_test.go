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
		deviceData  *DeviceData
		x           *float64
		y           *float64
		z           *float64
		expectError bool
	}{
		{
			name:        "Valid Device and Gyroscope Data",
			deviceData:  d,
			x:           floatPtr(1.0),
			y:           floatPtr(0.5),
			z:           floatPtr(-0.5),
			expectError: false,
		},
		{
			name:        "Nil Device data",
			deviceData:  nil,
			x:           floatPtr(1.0),
			y:           floatPtr(0.5),
			z:           floatPtr(-0.5),
			expectError: true,
		},
		{
			name:        "Missing Y value",
			deviceData:  d,
			x:           floatPtr(1.0),
			y:           nil,
			z:           floatPtr(-0.5),
			expectError: true,
		},
		{
			name:        "Missing Z value",
			deviceData:  d,
			x:           floatPtr(1.0),
			y:           floatPtr(0.5),
			z:           nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gyroData, err := NewGyroscope(tt.deviceData, tt.x, tt.y, tt.z)

			if tt.expectError {
				if err == nil {
					t.Error("Expected an error but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Did not expect an error but got %v", err)
				}
				if gyroData.DeviceData != tt.deviceData {
					t.Errorf("Expected Device %v, got %v", tt.deviceData, gyroData.DeviceData)
				}
				if gyroData.X != nil && *gyroData.X != *tt.x {
					t.Errorf("Expected X %f, got %f", *tt.x, *gyroData.X)
				}
				if gyroData.Y != nil && *gyroData.Y != *tt.y {
					t.Errorf("Expected Y %f, got %f", *tt.y, *gyroData.Y)
				}
				if gyroData.Z != nil && *gyroData.Z != *tt.z {
					t.Errorf("Expected Z %f, got %f", *tt.z, *gyroData.Z)
				}
			}
		})
	}
}

func floatPtr(f float64) *float64 {
	return &f
}
