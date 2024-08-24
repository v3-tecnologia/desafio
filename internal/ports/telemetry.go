package ports

import (
	"github.com/ThalesMonteir0/desafio/internal/domain"
	"github.com/ThalesMonteir0/desafio/pkg/err_rest"
)

type TelemetryRepository interface {
	CreateGyroscope(gyroscopeDomain domain.GyroscopeDomain) *err_rest.ErrRest
	CreateGps(gpsDomain domain.GpsDomain) *err_rest.ErrRest
	CreatePhoto(photoDomain domain.PhotoDomain) *err_rest.ErrRest
	FindDeviceByMAC(macAddress string) domain.DeviceDomain
}
