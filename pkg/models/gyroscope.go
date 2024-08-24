package models

type Gyroscope struct {
	*DeviceData
	X, Y, Z float64
}
