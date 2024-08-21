package mapper

import (
	"time"

	"github.com/charmingruby/g3/internal/telemetry/domain/entity"
)

func DomainGyroscopeToPostgres(gyroscope entity.Gyroscope) PostgresGyroscope {
	return PostgresGyroscope{
		ID:        gyroscope.ID,
		XPosition: gyroscope.XPosition,
		YPosition: gyroscope.YPosition,
		ZPosition: gyroscope.ZPosition,
		CreatedAt: gyroscope.CreatedAt,
	}
}

func PostgresGyroscopeToDomain(gyroscope PostgresGyroscope) entity.Gyroscope {
	return entity.Gyroscope{
		ID:        gyroscope.ID,
		XPosition: gyroscope.XPosition,
		YPosition: gyroscope.YPosition,
		ZPosition: gyroscope.ZPosition,
		CreatedAt: gyroscope.CreatedAt,
	}
}

type PostgresGyroscope struct {
	ID        string    `json:"id" db:"id"`
	XPosition float64   `json:"x" db:"x_position"`
	YPosition float64   `json:"y" db:"y_position"`
	ZPosition float64   `json:"z" db:"z_position"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
