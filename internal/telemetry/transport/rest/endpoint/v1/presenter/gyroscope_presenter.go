package presenter

import (
	"time"

	"github.com/charmingruby/g3/internal/telemetry/domain/entity"
)

func DomainGyroscopeToHTTP(gyroscope entity.Gyroscope) HTTPGyroscope {
	return HTTPGyroscope{
		ID:        gyroscope.ID,
		XPosition: gyroscope.XPosition,
		YPosition: gyroscope.YPosition,
		ZPosition: gyroscope.ZPosition,
		CreatedAt: gyroscope.CreatedAt,
	}

}

type HTTPGyroscope struct {
	ID        string    `json:"id"`
	XPosition float64   `json:"x"`
	YPosition float64   `json:"y"`
	ZPosition float64   `json:"z"`
	CreatedAt time.Time `json:"created_at"`
}
