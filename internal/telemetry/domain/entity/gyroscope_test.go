package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewGyroscope(t *testing.T) {
	t.Run("it should be able to assign all fields", func(t *testing.T) {
		x := 120.99
		y := -39.99
		z := -400.41

		g, err := NewGyroscope(
			GyroscopeProps{
				XPosition: x,
				YPosition: y,
				ZPosition: z,
			},
		)

		assert.NoError(t, err)
		assert.NotNil(t, g)
		assert.NotEmpty(t, g.ID)
		assert.Equal(t, x, g.XPosition)
		assert.Equal(t, y, g.YPosition)
		assert.Equal(t, z, g.ZPosition)
	})
}
