package entity

import (
	"time"

	"github.com/charmingruby/g3/internal/common/core"
)

type GyroscopeProps struct {
	XPosition float64
	YPosition float64
	ZPosition float64
}

func NewGyroscope(props GyroscopeProps) (*Gyroscope, error) {
	g := Gyroscope{
		ID:        core.NewID(),
		XPosition: props.XPosition,
		YPosition: props.YPosition,
		ZPosition: props.ZPosition,
		CreatedAt: time.Now(),
	}

	// como nao tem validaçoes a se fazer, caso precise, criar o metodo depois

	return &g, nil
}

type Gyroscope struct {
	ID        string
	XPosition float64
	YPosition float64
	ZPosition float64
	CreatedAt time.Time
}
