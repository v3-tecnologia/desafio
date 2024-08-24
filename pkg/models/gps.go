package models

type GPS struct {
	*DeviceData
	Latitude, Longitude float64
}
