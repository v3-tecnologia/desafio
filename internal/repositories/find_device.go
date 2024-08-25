package repositories

import (
	"github.com/ThalesMonteir0/desafio/internal/domain"
	"github.com/ThalesMonteir0/desafio/pkg/err_rest"
)

func (t telemetryRepository) FindDeviceByMAC(macAddress string) (domain.DeviceDomain, *err_rest.ErrRest) {
	//TODO implement me
	panic("implement me")
}
