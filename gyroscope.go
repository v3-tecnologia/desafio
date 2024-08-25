package main

import (
	"database/sql"
	"encoding/json"
)

type gyroscope struct {
	x, y, z   float64
	timestamp uint64
	deviceID  string
}

func (g *gyroscope) decode(data []byte) bool {
	var m map[string]interface{}
	var valid_device, valid_x, valid_y, valid_z, valid_time bool

	err := json.Unmarshal(data, &m)
	if err != nil {
		return false
	}

	g.deviceID, valid_device = validateDevice(m)
	g.timestamp, valid_time = validateTimestamp(m)
	g.x, valid_x = validateX(m)
	g.y, valid_y = validateY(m)
	g.z, valid_z = validateZ(m)

	return valid_device && valid_time && valid_x && valid_y && valid_z
}

func validateX(vals map[string]interface{}) (float64, bool) {
	x, valid := vals["x"].(float64)
	return x, valid
}

func validateY(vals map[string]interface{}) (float64, bool) {
	y, valid := vals["y"].(float64)
	return y, valid
}

func validateZ(vals map[string]interface{}) (float64, bool) {
	z, valid := vals["z"].(float64)
	return z, valid
}

func (g *gyroscope) persist(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO gyroscope (deviceID, x, y, z, time) VALUES (?, ?, ?, ?, ?)", g.deviceID, g.x, g.y, g.z, g.timestamp)
	return err
}
