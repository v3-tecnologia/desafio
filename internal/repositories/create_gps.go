package repositories

import (
	"github.com/ThalesMonteir0/desafio/internal/domain"
	"github.com/ThalesMonteir0/desafio/internal/repositories/convert"
	"github.com/ThalesMonteir0/desafio/pkg/err_rest"
)

func (t telemetryRepository) CreateGps(gpsDomain domain.GpsDomain) *err_rest.ErrRest {
	gpsEntity := convert.ConvertGpsDomainToEntity(gpsDomain)
	sql := `INSERT INTO public.gps(
	device_id, latitude, longitude, collection_date)
	VALUES ($1, $2, $3, $4);`

	result, err := t.db.Exec(sql, gpsEntity.DeviceID, gpsEntity.Latitude, gpsEntity.Longitude, gpsEntity.CollectionDate)
	if err != nil {
		return err_rest.NewInternalServerError("unable create gps")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return err_rest.NewInternalServerError("unable create gps")
	}

	return nil
}
