package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewPhoto(t *testing.T) {
	t.Run("it should be able to assign all fields", func(t *testing.T) {
		imageURL := "image_url"
		isRecognized := true

		p := NewPhoto(imageURL, isRecognized)

		assert.NotEmpty(t, p.ID)
		assert.Equal(t, imageURL, p.ImageURL)
		assert.Equal(t, isRecognized, p.IsRecognized)
	})
}
