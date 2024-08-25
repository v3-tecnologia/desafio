package convert

import (
	"github.com/ThalesMonteir0/desafio/internal/domain"
	"github.com/ThalesMonteir0/desafio/internal/repositories/entity"
)

func ConvertPhotoDomainToEntity(photoDomain domain.PhotoDomain) entity.Photos {
	return entity.Photos{
		ID:             photoDomain.ID,
		Url:            photoDomain.Url,
		DeviceID:       photoDomain.DeviceID,
		CollectionDate: &photoDomain.CollectionDate,
	}
}
