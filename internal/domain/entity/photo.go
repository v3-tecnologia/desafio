package entity

import (
	"errors"
	"github.com/google/uuid"
)

type Photo struct {
	id       string
	filePath string
	Device
}

func NewPhoto(filePath string, macAddress string) (*Photo, error) {
	device, err := NewDevice(macAddress)
	if err != nil {
		return nil, err
	}
	photo := &Photo{
		id:       uuid.New().String(),
		filePath: filePath,
		Device:   *device,
	}
	err = photo.IsValid()
	if err != nil {
		return nil, err
	}
	return photo, nil
}

func (p *Photo) IsValid() error {
	if p.id == "" {
		return errors.New("invalid id")
	}
	if p.filePath == "" {
		return errors.New("file path is required")
	}
	return nil
}

func (p *Photo) GetID() string {
	return p.id
}

func (p *Photo) GetFilePath() string {
	return p.filePath
}
