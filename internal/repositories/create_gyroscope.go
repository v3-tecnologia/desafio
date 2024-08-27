package repositories

import (
	"github.com/ThalesMonteir0/desafio/internal/domain"
	"github.com/ThalesMonteir0/desafio/internal/repositories/convert"
	"github.com/ThalesMonteir0/desafio/pkg/err_rest"
)

func (t telemetryRepository) CreateGyroscope(gyroscopeDomain domain.GyroscopeDomain) *err_rest.ErrRest {
	entity := convert.ConvertGyroscopeDomainToEntity(gyroscopeDomain)
	sql := `INSERT INTO public.gyroscopes(
		device_id, x, y, z, collection_date)
		VALUES ($1, $2, $3, $4, $5);`

	result, err := t.db.Exec(sql, entity.DeviceID, entity.X, entity.Y, entity.Z, entity.CollectionDate)
	if err != nil {
		return err_rest.NewInternalServerError("unable create gyroscope")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return err_rest.NewInternalServerError("unable create gyroscope")
	}

	return nil
}
