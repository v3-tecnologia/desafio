package service

import (
	"github.com/ThalesMonteir0/desafio/internal/domain"
	"github.com/ThalesMonteir0/desafio/pkg/err_rest"
)

func (t telemetryService) findDevice(macAddress string) (domain.DeviceDomain, *err_rest.ErrRest) {
	device, err := t.repository.FindDeviceByMAC(macAddress)
	if err != nil {
		return domain.DeviceDomain{}, err
	}

	return device, nil
}
