package dummy

import (
	"bytes"
	"github.com/kevenmiano/v3/internal/domain"
	"github.com/kevenmiano/v3/internal/shared"
	"mime/multipart"
)

func NewMultipart(f *domain.File) *shared.Request {

	var body bytes.Buffer

	writer := multipart.NewWriter(&body)

	part, _ := writer.CreateFormFile("file", f.GetName())

	if _, err := part.Write(f.GetContent()); err != nil {
		return nil
	}

	_ = writer.Close()

	return &shared.Request{
		Body: body.String(),
		Headers: map[string]string{
			"Content-Type": writer.FormDataContentType(),
		},
	}
}
