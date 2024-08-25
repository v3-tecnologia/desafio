package repositories

import (
	"github.com/ThalesMonteir0/desafio/internal/domain"
	"github.com/ThalesMonteir0/desafio/internal/repositories/convert"
	"github.com/ThalesMonteir0/desafio/pkg/err_rest"
)

func (t telemetryRepository) CreatePhoto(photoDomain domain.PhotoDomain) *err_rest.ErrRest {
	photoEntity := convert.ConvertPhotoDomainToEntity(photoDomain)
	sql := `INSERT INTO public.photos(
	device_id, url, collection_date)
	VALUES ($1, $2, $3);`

	result, err := t.db.Exec(sql, photoEntity.DeviceID, photoEntity.Url, photoEntity.CollectionDate)
	if err != nil {
		return err_rest.NewInternalServerError("unable create photo")
	}

	if rowsAffected, err := result.RowsAffected(); err != nil || rowsAffected == 0 {
		return err_rest.NewInternalServerError("unable create photo")
	}

	return nil
}
