package domain

import (
	"errors"
	"mime/multipart"
)

type Reader struct {
	multipart.Reader
}

func NewReader(r *multipart.Reader) *Reader {

	return &Reader{
		Reader: *r,
	}

}

func (r *Reader) Validate() (bool, error) {

	if r == nil {
		return false, ErrInvalidReader
	}

	return true, nil
}

var (
	ErrInvalidReader = errors.New("invalid reader")
)
