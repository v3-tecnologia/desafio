package dummy

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/kevenmiano/v3/internal/domain"
)

func Gps() *domain.GPS {
	return domain.NewGPS(&domain.GPSDto{
		DeviceID: gofakeit.UUID(),
		CoordinateDto: domain.CoordinateDto{
			Latitude:  gofakeit.Latitude(),
			Longitude: gofakeit.Longitude(),
		},
	})
}
