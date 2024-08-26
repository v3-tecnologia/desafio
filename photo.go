package main

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"os"
	"strconv"
)

type photo struct {
	image     []byte
	file      string
	format    string
	timestamp uint64
	deviceID  string
}

const photo_dir = `./`

func (p *photo) decode(data []byte) bool {
	var m map[string]interface{}
	var valid_device, valid_image, valid_time, valid_format bool

	err := json.Unmarshal(data, &m)
	if err != nil {
		return false
	}

	p.deviceID, valid_device = validateDevice(m)
	p.timestamp, valid_time = validateTimestamp(m)
	p.image, valid_image = validateImage(m)
	p.format, valid_format = validateFormat(m)

	valid_file := valid_device && valid_time && valid_image && valid_format
	if valid_file {
		p.file = photo_dir + p.deviceID + "-" + strconv.FormatUint(p.timestamp, 10) + "." + p.format
	}

	return valid_file
}

func validateFormat(vals map[string]interface{}) (string, bool) {
	format, valid := vals["format"].(string)

	if valid {
		switch format {
		case "jpg":
		case "png":
		case "gif":
		case "bmp":
		default:
			valid = false
		}
	}

	return format, valid
}

func validateImage(vals map[string]interface{}) ([]byte, bool) {
	var image []byte
	var err error
	image_str, valid := vals["image"].(string)

	if valid {
		image, err = base64.StdEncoding.DecodeString(image_str)
		valid = (err == nil)
	}

	return image, valid
}

func (p *photo) persist(db *sql.DB) error {
	file, err := os.Create(p.file)
	if err != nil {
		return err
	}

	_, err = file.Write(p.image)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO photos (deviceID, photo, time) VALUES (?, ?, ?)", p.deviceID, p.file, p.timestamp)
	return err
}
