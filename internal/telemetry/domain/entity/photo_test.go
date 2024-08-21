package entity

import (
	"testing"

	"github.com/charmingruby/g3/internal/common/custom_err"
	"github.com/stretchr/testify/assert"
)

func Test_NewPhoto(t *testing.T) {
	t.Run("it should be able to assign all fields if valid", func(t *testing.T) {
		imageURL := "image_url"
		isRecognized := true

		p, err := NewPhoto(
			PhotoProps{
				ImageURL:     imageURL,
				IsRecognized: isRecognized,
			},
		)

		assert.NotNil(t, p)
		assert.NoError(t, err)
		assert.NotEmpty(t, p.ID)
		assert.Equal(t, imageURL, p.ImageURL)
		assert.Equal(t, isRecognized, p.IsRecognized)
	})

	t.Run("it should be not able to assign all fields if params are invalid", func(t *testing.T) {
		isRecognized := true

		p, err := NewPhoto(
			PhotoProps{
				IsRecognized: isRecognized,
			},
		)

		assert.Nil(t, p)
		assert.Error(t, err)
		assert.Equal(t,
			custom_err.NewValidationErr(custom_err.NewRequiredErrMessage("image_url")).Error(),
			err.Error(),
		)
	})
}
