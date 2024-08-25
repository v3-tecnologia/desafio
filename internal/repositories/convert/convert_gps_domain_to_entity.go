package convert

import (
	"github.com/ThalesMonteir0/desafio/internal/domain"
	"github.com/ThalesMonteir0/desafio/internal/repositories/entity"
)

func ConvertGpsDomainToEntity(gpsDomain domain.GpsDomain) entity.Gps {
	return entity.Gps{
		ID:             gpsDomain.ID,
		Latitude:       gpsDomain.Latitude,
		Longitude:      gpsDomain.Longitude,
		CollectionDate: &gpsDomain.CollectionDate,
		DeviceID:       gpsDomain.DeviceID,
	}
}
