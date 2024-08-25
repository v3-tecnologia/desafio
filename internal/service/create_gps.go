package service

import (
	"github.com/ThalesMonteir0/desafio/internal/domain"
	"github.com/ThalesMonteir0/desafio/pkg/err_rest"
)

func (t telemetryService) CreateGpsService(gpsDomain domain.GpsDomain) *err_rest.ErrRest {
	device, err := t.findDevice(gpsDomain.MacAddress)
	if err != nil {
		return err
	}

	gpsDomain.DeviceID = device.ID

	if restErr := t.repository.CreateGps(gpsDomain); restErr != nil {
		return restErr
	}

	return nil
}
