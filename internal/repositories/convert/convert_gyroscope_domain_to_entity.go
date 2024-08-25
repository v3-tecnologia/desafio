package convert

import (
	"github.com/ThalesMonteir0/desafio/internal/domain"
	"github.com/ThalesMonteir0/desafio/internal/repositories/entity"
)

func ConvertGyroscopeDomainToEntity(gyroscopeDomain domain.GyroscopeDomain) entity.Gyroscopes {
	return entity.Gyroscopes{
		ID:             gyroscopeDomain.ID,
		X:              gyroscopeDomain.X,
		Y:              gyroscopeDomain.Y,
		Z:              gyroscopeDomain.Z,
		DeviceID:       gyroscopeDomain.DeviceID,
		CollectionDate: &gyroscopeDomain.CollectionDate,
	}
}
