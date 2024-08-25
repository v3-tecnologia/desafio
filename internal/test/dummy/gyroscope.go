package dummy

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/kevenmiano/v3/internal/domain"
)

func Gyroscope() *domain.Gyroscope {
	return domain.NewGyroscope(&domain.GyroscopeDto{
		DeviceID: gofakeit.UUID(),
		X:        gofakeit.Float64(),
		Y:        gofakeit.Float64(),
		Z:        gofakeit.Float64(),
	})

}
