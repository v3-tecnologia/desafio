package service

import (
	"github.com/ThalesMonteir0/desafio/internal/domain"
	"github.com/ThalesMonteir0/desafio/pkg/err_rest"
)

func (t telemetryService) CreateGyroscopeService(gyroscopeDomain domain.GyroscopeDomain) *err_rest.ErrRest {
	device, err := t.findDevice(gyroscopeDomain.MacAddress)
	if err != nil {
		return err
	}

	gyroscopeDomain.DeviceID = device.ID

	if err = t.repository.CreateGyroscope(gyroscopeDomain); err != nil {
		return err
	}

	return nil
}
