package util

import (
	"fmt"
	"math/rand"
)

const (
	local     = 0b10
	multicast = 0b1
)

func GenerateMac() string {
	buf := make([]byte, 6)
	_, err := rand.Read(buf)
	if err != nil {
		return ""
	}
	buf[0] = buf[0]&^multicast | local

	m := fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", buf[0], buf[1], buf[2], buf[3], buf[4], buf[5])
	return m
}
