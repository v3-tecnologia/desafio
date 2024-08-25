package convert

import (
	"github.com/ThalesMonteir0/desafio/internal/domain"
	"github.com/ThalesMonteir0/desafio/internal/repositories/entity"
)

func ConvertDeviceEntityToDomain(devices entity.Devices) domain.DeviceDomain {
	return domain.DeviceDomain{
		ID:  devices.ID,
		Mac: devices.MacAddress,
	}
}
