package form

import (
	"errors"
	"github.com/grokify/go-awslambda"
	"github.com/kevenmiano/v3/internal/domain"
	"github.com/kevenmiano/v3/internal/shared"
	"io"
	"path/filepath"
)

type Data struct{}

func NewFormData() *Data {
	return &Data{}
}

func (m *Data) Read(request shared.Request) (*domain.File, error) {

	reader, err := m.reader(request)

	if err != nil {
		return nil, err
	}

	return reader, nil
}

func (m *Data) reader(r shared.Request) (*domain.File, error) {

	multipart, err := awslambda.NewReaderMultipart(r)

	if err != nil {
		return nil, ErrInvalidFormReader

	}

	reader := domain.NewReader(multipart)

	if ok, e := reader.Validate(); !ok {
		return nil, e
	}

	part, err := reader.NextPart()

	if err != nil {
		return nil, ErrInvalidFormFile
	}

	content, err := io.ReadAll(part)

	if err != nil {
		return nil, ErrReadFormFile
	}

	maxSize := 1024 * 1024 * 5

	if len(content) > maxSize {
		return nil, ErrorFileTooLarge
	}

	fileName := part.FileName()

	file := domain.NewFile(&domain.FileDto{
		Name:    fileName,
		Content: content,
		Ext:     filepath.Ext(fileName),
	})

	return file, nil
}

var (
	ErrInvalidFormReader = errors.New("file not found, attach a file with the key 'file'")
	ErrInvalidFormFile   = errors.New("invalid file")
	ErrReadFormFile      = errors.New("error reading file")
	ErrorFileTooLarge    = errors.New("file too large")
)
