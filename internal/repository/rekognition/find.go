package rekognition

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/kevenmiano/v3/internal/domain"
)

func (r *IRepository) Find(d *domain.Faces) (bool, error) {

	searchFacesInput := &rekognition.SearchFacesByImageInput{
		CollectionId: aws.String(os.Getenv("REKOGNITION_COLLECTION")),
		Image: &types.Image{
			S3Object: &types.S3Object{
				Bucket: aws.String(os.Getenv("V3_BUCKET")),
				Name:   aws.String(fmt.Sprintf("uploads/%s", d.ObjectKey)),
			},
		},
	}

	result, err := r.database.SearchFacesByImage(context.Background(), searchFacesInput)

	if err != nil {
		return false, err
	}

	if len(result.FaceMatches) == 0 {
		return false, nil
	}

	return true, nil

}
