package domain

import "github.com/HaroldoFV/desafio/internal/domain/entity"

type GyroscopeRepositoryInterface interface {
	Create(gyroscope *entity.Gyroscope) error
}
