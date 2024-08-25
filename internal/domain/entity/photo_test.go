package entity

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewPhoto(t *testing.T) {
	t.Run("should create a new photo successfully", func(t *testing.T) {
		filePath := "/path/to/image.jpg"
		macAddress := "00:11:22:33:44:55"

		photo, err := NewPhoto(filePath, macAddress)

		assert.NoError(t, err)
		assert.NotNil(t, photo)
		assert.NotEmpty(t, photo.GetID())
		assert.Equal(t, filePath, photo.GetFilePath())
		assert.Equal(t, macAddress, photo.GetMACAddress())
	})

	t.Run("should return error when file path is empty", func(t *testing.T) {
		filePath := ""
		macAddress := "00:11:22:33:44:55"

		photo, err := NewPhoto(filePath, macAddress)

		assert.Error(t, err)
		assert.Nil(t, photo)
		assert.EqualError(t, err, "file path is required")
	})

	t.Run("should return error when mac address is invalid", func(t *testing.T) {
		filePath := "/path/to/image.jpg"
		macAddress := "invalid mac"

		photo, err := NewPhoto(filePath, macAddress)

		assert.Error(t, err)
		assert.Nil(t, photo)
		assert.EqualError(t, err, "invalid MAC address format")
	})
}

func TestPhoto_GetMethods(t *testing.T) {
	filePath := "/path/to/image.jpg"
	macAddress := "00:11:22:33:44:55"

	photo, _ := NewPhoto(filePath, macAddress)

	t.Run("GetID", func(t *testing.T) {
		require.NotEmpty(t, photo.GetID())
	})

	t.Run("GetFilePath", func(t *testing.T) {
		require.Equal(t, "/path/to/image.jpg", photo.GetFilePath())
	})

	t.Run("GetMACAddress", func(t *testing.T) {
		require.Equal(t, "00:11:22:33:44:55", photo.GetMACAddress())
	})

	t.Run("GetTimestamp", func(t *testing.T) {
		require.NotZero(t, photo.GetTimestamp())
	})
}
