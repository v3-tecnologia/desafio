package entity

import (
	"errors"
	"regexp"
	"time"
)

type Device struct {
	macAddress string
	timestamp  time.Time
}

func NewDevice(macAddress string) (*Device, error) {
	if !isValidMACAddress(macAddress) {
		return nil, errors.New("invalid MAC address format")
	}
	return &Device{
		macAddress: macAddress,
		timestamp:  time.Now(),
	}, nil
}

func isValidMACAddress(mac string) bool {
	match, _ := regexp.MatchString(`^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$`, mac)
	return match
}

func (g *Device) GetMACAddress() string {
	return g.macAddress
}

func (g *Device) GetTimestamp() time.Time {
	return g.timestamp
}
