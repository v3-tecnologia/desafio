package dummy

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/kevenmiano/v3/internal/domain"
)

func Photo() *domain.Photo {
	return domain.NewPhoto(&domain.PhotoDto{
		FileName:    fmt.Sprintf("%s.jpg", gofakeit.UUID()),
		Content:     gofakeit.ImagePng(200, 200),
		ContentType: ".png",
	})
}
