package entity

import (
	"time"

	"github.com/charmingruby/g3/internal/common/core"
)

type GPSProps struct {
	Latitude  float64
	Longitude float64
}

func NewGPS(props GPSProps) (*GPS, error) {
	g := GPS{
		ID:        core.NewID(),
		Latitude:  props.Latitude,
		Longitude: props.Longitude,
		CreatedAt: time.Now(),
	}

	// como nao tem validaçoes a se fazer, caso precise, criar o metodo depois

	return &g, nil
}

type GPS struct {
	ID        string
	Latitude  float64
	Longitude float64
	CreatedAt time.Time
}
