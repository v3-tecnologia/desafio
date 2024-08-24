package entity

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewGPS(t *testing.T) {
	t.Run("Valid GPS", func(t *testing.T) {
		gps, err := NewGPS(40.7128, -74.0060, "00:11:22:33:44:55")
		require.Nil(t, err)
		require.NotNil(t, gps)
		require.Equal(t, 40.7128, gps.GetLatitude())
		require.Equal(t, -74.0060, gps.GetLongitude())
		require.Equal(t, "00:11:22:33:44:55", gps.GetMACAddress())
	})

	t.Run("Invalid Latitude", func(t *testing.T) {
		_, err := NewGPS(91.0, -74.0060, "00:11:22:33:44:55")
		require.EqualError(t, err, "latitude must be between -90 and 90")
	})

	t.Run("Invalid Longitude", func(t *testing.T) {
		_, err := NewGPS(40.7128, 181.0, "00:11:22:33:44:55")
		require.EqualError(t, err, "longitude must be between -180 and 180")
	})

	t.Run("Invalid MAC Address", func(t *testing.T) {
		_, err := NewGPS(40.7128, -74.0060, "invalid-mac")
		require.EqualError(t, err, "invalid MAC address format")
	})
}

func TestGPS_GetMethods(t *testing.T) {
	gps, _ := NewGPS(40.7128, -74.0060, "00:1A:2B:3C:4D:5E")

	t.Run("GetID", func(t *testing.T) {
		require.NotEmpty(t, gps.GetID())
	})

	t.Run("GetLatitude", func(t *testing.T) {
		require.Equal(t, 40.7128, gps.GetLatitude())
	})

	t.Run("GetLongitude", func(t *testing.T) {
		require.Equal(t, -74.0060, gps.GetLongitude())
	})

	t.Run("GetMACAddress", func(t *testing.T) {
		require.Equal(t, "00:1A:2B:3C:4D:5E", gps.GetMACAddress())
	})

	t.Run("GetTimestamp", func(t *testing.T) {
		require.NotZero(t, gps.GetTimestamp())
	})
}

func TestGPS_IsValid(t *testing.T) {
	t.Run("Valid GPS", func(t *testing.T) {
		gps, _ := NewGPS(40.7128, -74.0060, "00:11:22:33:44:55")
		err := gps.IsValid()
		require.Nil(t, err)
	})

	t.Run("Invalid ID", func(t *testing.T) {
		gps, _ := NewGPS(40.7128, -74.0060, "00:11:22:33:44:55")
		gps.id = ""
		err := gps.IsValid()
		require.EqualError(t, err, "invalid id")
	})

	t.Run("Invalid Latitude", func(t *testing.T) {
		gps, _ := NewGPS(40.7128, -74.0060, "00:11:22:33:44:55")
		gps.latitude = 91.0
		err := gps.IsValid()
		require.EqualError(t, err, "latitude must be between -90 and 90")
	})

	t.Run("Invalid Longitude", func(t *testing.T) {
		gps, _ := NewGPS(40.7128, -74.0060, "00:11:22:33:44:55")
		gps.longitude = 181.0
		err := gps.IsValid()
		require.EqualError(t, err, "longitude must be between -180 and 180")
	})
}
