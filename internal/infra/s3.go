package infra

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3 struct {
	session *session.Session
}

type NewS3Client interface {
	Put(body *s3.PutObjectInput) (*s3.PutObjectOutput, error)
}

func NewS3() *S3 {

	newSession, err := session.NewSession(&aws.Config{})

	if err != nil {
		_ = fmt.Errorf("error creating new session: %v", err)
		panic(err)
	}

	return &S3{
		session: newSession,
	}
}

func (s *S3) Put(body *s3.PutObjectInput) (*s3.PutObjectOutput, error) {

	svc := s3.New(s.session)

	res, err := svc.PutObject(body)

	if err != nil {
		_ = fmt.Errorf("error uploading to s3: %v", err)
		return nil, err
	}

	return res, nil

}
