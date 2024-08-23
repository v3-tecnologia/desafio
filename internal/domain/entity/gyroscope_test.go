package entity

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewGyroscope(t *testing.T) {
	t.Run("Valid Gyroscope", func(t *testing.T) {
		gyro, err := NewGyroscope("MPU-6050", "InvenSense GY-521",
			1.0, 2.0, 3.0, "00:11:22:33:44:55")
		require.Nil(t, err)
		require.NotNil(t, gyro)
		require.Equal(t, "MPU-6050", gyro.GetName())
		require.Equal(t, "InvenSense GY-521", gyro.GetModel())
		require.Equal(t, "00:11:22:33:44:55", gyro.GetMACAddress())
	})

	t.Run("Invalid Name", func(t *testing.T) {
		_, err := NewGyroscope("", "InvenSense GY-521",
			1.0, 2.0, 3.0, "00:11:22:33:44:55")
		require.EqualError(t, err, "name cannot be empty")
	})

	t.Run("Name Too Long", func(t *testing.T) {
		longName := string(make([]byte, 101))
		_, err := NewGyroscope(longName, "InvenSense GY-521",
			1.0, 2.0, 3.0, "00:11:22:33:44:55")
		require.EqualError(t, err, "name cannot be longer than 100 characters")
	})

	t.Run("Invalid Model", func(t *testing.T) {
		_, err := NewGyroscope("MPU-6050", "", 1.0, 2.0, 3.0,
			"00:11:22:33:44:55")
		require.EqualError(t, err, "model cannot be empty")
	})

	t.Run("Model Too Long", func(t *testing.T) {
		longModel := string(make([]byte, 51))
		_, err := NewGyroscope("MPU-6050", longModel, 1.0, 2.0, 3.0,
			"00:11:22:33:44:55")
		require.EqualError(t, err, "model cannot be longer than 50 characters")
	})

	t.Run("Invalid MAC Address", func(t *testing.T) {
		_, err := NewGyroscope("MPU-6050", "InvenSense GY-521",
			1.0, 2.0, 3.0, "invalid-mac")
		require.EqualError(t, err, "invalid MAC address format")
	})

}

func TestGyroscope_GetMethods(t *testing.T) {
	gyro, _ := NewGyroscope("MPU-6050", "InvenSense GY-521",
		1.0, 2.0, 3.0, "00:1A:2B:3C:4D:5E")

	t.Run("GetID", func(t *testing.T) {
		require.NotEmpty(t, gyro.GetID())
	})

	t.Run("GetName", func(t *testing.T) {
		require.Equal(t, "MPU-6050", gyro.GetName())
	})

	t.Run("GetModel", func(t *testing.T) {
		require.Equal(t, "InvenSense GY-521", gyro.GetModel())
	})

	t.Run("GetMACAddress", func(t *testing.T) {
		require.Equal(t, "00:1A:2B:3C:4D:5E", gyro.GetMACAddress())
	})

	t.Run("GetX", func(t *testing.T) {
		require.Equal(t, 1.0, gyro.GetX())
	})

	t.Run("GetY", func(t *testing.T) {
		require.Equal(t, 2.0, gyro.GetY())
	})

	t.Run("GetZ", func(t *testing.T) {
		require.Equal(t, 3.0, gyro.GetZ())
	})

	t.Run("GetTimestamp", func(t *testing.T) {
		require.NotZero(t, gyro.GetTimestamp())
	})
}

func TestGyroscope_Update(t *testing.T) {
	t.Run("Update with valid data", func(t *testing.T) {
		gyro, _ := NewGyroscope("MPU-6050", "InvenSense GY-521",
			1.0, 2.0, 3.0, "00:11:22:33:44:55")
		err := gyro.Update("L3GD20H", "STMicroelectronics", "00:AA:BB:CC:DD:EE")
		require.Nil(t, err)
		require.Equal(t, "L3GD20H", gyro.GetName())
		require.Equal(t, "STMicroelectronics", gyro.GetModel())
		require.Equal(t, "00:AA:BB:CC:DD:EE", gyro.GetMACAddress())
	})

	t.Run("Update with empty name", func(t *testing.T) {
		gyro, _ := NewGyroscope("MPU-6050", "InvenSense GY-521",
			1.0, 2.0, 3.0, "00:11:22:33:44:55")
		err := gyro.Update("", "STMicroelectronics", "00:AA:BB:CC:DD:EE")
		require.EqualError(t, err, "name cannot be empty")
	})

	t.Run("Update with too long name", func(t *testing.T) {
		gyro, _ := NewGyroscope("MPU-6050", "InvenSense GY-521",
			1.0, 2.0, 3.0, "00:11:22:33:44:55")
		longName := string(make([]byte, 101))
		err := gyro.Update(longName, "STMicroelectronics", "00:AA:BB:CC:DD:EE")
		require.EqualError(t, err, "name cannot be longer than 100 characters")
	})

	t.Run("Update with invalid MAC address", func(t *testing.T) {
		gyro, _ := NewGyroscope("MPU-6050", "InvenSense GY-521",
			1.0, 2.0, 3.0, "00:11:22:33:44:55")
		err := gyro.Update("L3GD20H", "STMicroelectronics", "invalid-mac")
		require.EqualError(t, err, "invalid MAC address format")
	})
}
