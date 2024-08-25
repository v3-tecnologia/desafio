package models

import "errors"

type PhotoData struct {
	*DeviceData `json:"deviceData"`
	Path        string `json:"path"`
}

func NewPhotoData(d *DeviceData, p string) (*PhotoData, error) {
	if d == nil {
		return nil, errors.New("device data cannot be nil")
	}
	if len(p) == 0 {
		return nil, errors.New("path to photo cannot be empty")
	}

	return &PhotoData{
		DeviceData: d,
		Path:       p,
	}, nil
}
