package models

import (
	"testing"
	"time"
)

func TestNewPhotoData(t *testing.T) {
	validDevice := &DeviceData{
		MAC:       "AA:BB:CC:DD:EE:FF",
		Timestamp: time.Now(),
	}

	tests := []struct {
		name        string
		device      *DeviceData
		path        string
		expectError bool
	}{
		{
			name:        "Valid Device data and path to the photo",
			device:      validDevice,
			path:        "/photos/image.jpg",
			expectError: false,
		},
		{
			name:        "Nil Device",
			device:      nil,
			path:        "/photos/image.jpg",
			expectError: true,
		},
		{
			name:        "Empty path to the photo",
			device:      validDevice,
			path:        "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			photoData, err := NewPhotoData(tt.device, tt.path)

			if tt.expectError {
				if err == nil {
					t.Error("Expected an error but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Did not expect an error but got %v", err)
				}
				if photoData.DeviceData != tt.device {
					t.Errorf("Expected Device %v, got %v", tt.device, photoData.DeviceData)
				}
				if photoData.Path != tt.path {
					t.Errorf("Expected Path %s, got %s", tt.path, photoData.Path)
				}
			}
		})
	}
}
