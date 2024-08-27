package util

import (
	"errors"
	"net"
)

// IsValidateMacAddress checks to see if a MAC address string is well-formatted
func IsValidateMacAddress(mac string) error {
	_, err := net.ParseMAC(mac)
	if err != nil {
		return errors.New("Invalid MAC address: " + mac)
	}
	return nil
}
