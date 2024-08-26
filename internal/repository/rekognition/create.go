package rekognition

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/kevenmiano/v3/internal/domain"
)

func (r *IRepository) Create(d *domain.Faces) (*domain.Faces, error) {

	indexFaces := &rekognition.IndexFacesInput{
		CollectionId: aws.String(os.Getenv("REKOGNITION_COLLECTION")),
		Image: &types.Image{
			S3Object: &types.S3Object{
				Bucket: aws.String(os.Getenv("V3_BUCKET")),
				Name:   aws.String(d.ObjectKey),
			},
		},
	}

	_, err := r.database.IndexFaces(context.Background(), indexFaces)

	if err != nil {
		return nil, err
	}

	return d, nil
}
