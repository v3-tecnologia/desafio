package entity

import (
	"errors"
	"github.com/google/uuid"
	"regexp"
	"time"
)

type Gyroscope struct {
	id         string
	name       string
	model      string
	x          float64
	y          float64
	z          float64
	timestamp  time.Time
	macAddress string
}

func NewGyroscope(name, model, macAddress string) (*Gyroscope, error) {
	gyroscope := &Gyroscope{
		id:         uuid.New().String(),
		name:       name,
		model:      model,
		macAddress: macAddress,
		x:          0,
		y:          0,
		z:          0,
		timestamp:  time.Now(),
	}
	err := gyroscope.IsValid()
	if err != nil {
		return nil, err
	}
	return gyroscope, nil
}

func (g *Gyroscope) Update(name, model, macAddress string) error {
	g.name = name
	g.model = model
	g.macAddress = macAddress
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
	if !isValidMACAddress(g.macAddress) {
		return errors.New("invalid MAC address format")
	}
	return nil
}

func isValidMACAddress(mac string) bool {
	match, _ := regexp.MatchString(`^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$`, mac)
	return match
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

func (g *Gyroscope) GetMACAddress() string {
	return g.macAddress
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

func (g *Gyroscope) GetTimestamp() time.Time {
	return g.timestamp
}
