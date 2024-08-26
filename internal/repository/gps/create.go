package gps

import (
	"context"
	"github.com/kevenmiano/v3/internal/domain"
)

func (r *IRepository) Create(d *domain.GPS) (*domain.GPS, error) {

	err := r.database.Put(context.Background(), d)

	if err != nil {
		return nil, err
	}

	return d, nil
}
