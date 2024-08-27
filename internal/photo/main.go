package photo

import (
	"desafio-backend/pkg/errors"
	"mime/multipart"
)

type UseCases interface {
	ValidatePhoto(photo Request) errors.ErrorList
	ParseImage(file multipart.File, fileHeader *multipart.FileHeader) (ImageFile, errors.Error)
	ParsePhoto(photo string, file ImageFile) (Request, errors.Error)
	SavePhoto(photo Request) (Response, errors.Error)
}

type Main struct{}

func NewMain() UseCases {
	return Main{}
}
