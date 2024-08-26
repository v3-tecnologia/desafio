package service

import (
	"github.com/ThalesMonteir0/desafio/internal/domain"
	"github.com/ThalesMonteir0/desafio/pkg/err_rest"
)

func (t telemetryService) CreatePhotoService(photoDomain domain.PhotoDomain) *err_rest.ErrRest {
	device, err := t.repository.FindDeviceByMAC(photoDomain.MacAddress)
	if err != nil {
		return err
	}

	photoDomain.DeviceID = device.ID

	if err = t.repository.CreatePhoto(photoDomain); err != nil {
		return err
	}

	return nil
}
