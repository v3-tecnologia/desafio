package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/charmingruby/g3/config"
)

func NewAWSInstance(cfg config.Config) AWSInstance {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(
			cfg.AWSConfig.AccessKeyID,
			cfg.AWSConfig.SecretAccessKey,
			"",
		),
	}))

	rekognition := rekognition.New(sess)

	return AWSInstance{
		Session:     sess,
		Rekognition: rekognition,
	}
}

type AWSInstance struct {
	Session     *session.Session
	Rekognition *rekognition.Rekognition
}
