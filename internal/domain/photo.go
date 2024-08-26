package domain

import (
	"errors"
)

type Photo struct {
	ID          string `json:"id"`
	FileName    string `json:"fileName"`
	ContentType string `json:"type"`
	Content     []byte `json:"content"`
	Recognized  bool   `json:"recognized"`
}

type PhotoDto struct {
	FileName    string `json:"fileName"`
	ContentType string `json:"type"`
	Content     []byte `json:"content"`
}

func NewPhoto(d *PhotoDto) *Photo {
	return &Photo{
		FileName:    d.FileName,
		ContentType: d.ContentType,
		Content:     d.Content,
		Recognized:  false,
	}
}

func (f *Photo) GetKey() string {
	return f.ID + f.ContentType
}

func (f *Photo) Recognize() {
	f.Recognized = true
}

func (f *Photo) Validate() (bool, error) {
	if f.FileName == "" {
		return false, ErrFileNameRequiredPhoto
	}

	if len(f.Content) == 0 {
		return false, ErrContentPhoto
	}

	allowedMimeTypes := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}

	if _, ok := allowedMimeTypes[f.ContentType]; !ok {
		return false, ErrContentTypeInvalidPhoto
	}

	return true, nil
}

var (
	ErrFileNameRequiredPhoto   = errors.New("file name is required")
	ErrContentPhoto            = errors.New("content is required")
	ErrContentTypeInvalidPhoto = errors.New("content type is invalid, only .jpg/.jpeg/.png")
)
