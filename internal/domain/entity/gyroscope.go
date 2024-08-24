package entity

import (
	"errors"
	"github.com/google/uuid"
)

type Gyroscope struct {
	id    string
	name  string
	model string
	x     float64
	y     float64
	z     float64
	Device
}

func NewGyroscope(name, model string, x, y, z float64, macAddress string) (*Gyroscope, error) {
	device, err := NewDevice(macAddress)
	if err != nil {
		return nil, err
	}
	gyroscope := &Gyroscope{
		id:     uuid.New().String(),
		name:   name,
		model:  model,
		x:      x,
		y:      y,
		z:      z,
		Device: *device,
	}
	err = gyroscope.IsValid()
	if err != nil {
		return nil, err
	}
	return gyroscope, nil
}

func (g *Gyroscope) Update(name, model, macAddress string) error {
	device, err := NewDevice(macAddress)
	if err != nil {
		return err
	}
	g.name = name
	g.model = model
	g.Device = *device
	return g.IsValid()
}

func (g *Gyroscope) IsValid() error {
	if g.id == "" {
		return errors.New("invalid id")
	}
	if g.name == "" {
		return errors.New("name cannot be empty")
	}
	if len(g.name) > 100 {
		return errors.New("name cannot be longer than 100 characters")
	}
	if g.model == "" {
		return errors.New("model cannot be empty")
	}
	if len(g.model) > 50 {
		return errors.New("model cannot be longer than 50 characters")
	}
	if g.macAddress == "" {
		return errors.New("MAC address cannot be empty")
	}
	return nil
}

func (g *Gyroscope) GetID() string {
	return g.id
}

func (g *Gyroscope) GetName() string {
	return g.name
}

func (g *Gyroscope) GetModel() string {
	return g.model
}

func (g *Gyroscope) GetX() float64 {
	return g.x
}

func (g *Gyroscope) GetY() float64 {
	return g.y
}

func (g *Gyroscope) GetZ() float64 {
	return g.z
}
