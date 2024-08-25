package gyroscope

import (
	"fmt"

	"github.com/kevenmiano/v3/internal/domain"
)

func (l *Service) Create(d *domain.Gyroscope) (*domain.Gyroscope, error) {

	gyroscopeCreated, err := l.createGyroscopeUseCase.Execute(d)

	if err != nil {
		l.logger.Error(err.Error())
		return nil, err
	}

	l.logger.Info(fmt.Sprintf("Gyroscope created: %v", gyroscopeCreated))

	return gyroscopeCreated, nil

}
