package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewGPS(t *testing.T) {
	t.Run("it should be able to assign all fields", func(t *testing.T) {
		latitude := 419.99
		longitude := -194.99

		p, err := NewGPS(
			GPSProps{
				Latitude:  latitude,
				Longitude: longitude,
			},
		)

		assert.NoError(t, err)
		assert.NotNil(t, p)
		assert.NotEmpty(t, p.ID)
		assert.Equal(t, latitude, p.Latitude)
		assert.Equal(t, longitude, p.Longitude)
	})
}
