package s3

import (
	"bytes"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/kevenmiano/v3/internal/domain"
	"github.com/kevenmiano/v3/internal/infra"
)

func (s *IRepository) Upload(d *domain.Photo) (*domain.Photo, error) {

	svc := infra.NewS3()

	_, err := svc.Put(&s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("V3_BUCKET")),
		Key:    aws.String(fmt.Sprintf("uploads/%v%v", d.ID, d.ContentType)),
		Body:   bytes.NewReader(d.Content),
	})

	if err != nil {

		return nil, err
	}

	return d, nil
}
