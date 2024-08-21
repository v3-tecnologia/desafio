package adapter

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/charmingruby/g3/internal/common/constant"
	"github.com/charmingruby/g3/internal/telemetry/domain/port"
)

func NewAWSRekognitionRecognizerAdapter(r *rekognition.Rekognition) *AWSRekognitionRecognizerAdapter {
	return &AWSRekognitionRecognizerAdapter{
		rekognitionClient: r,
	}
}

type AWSRekognitionRecognizerAdapter struct {
	rekognitionClient *rekognition.Rekognition
}

func (r *AWSRekognitionRecognizerAdapter) Recognize(imageURL string) ([]port.DetectedFace, error) {
	detectedFaces := []port.DetectedFace{}

	imageBytes, err := os.ReadFile(constant.FilesDirectory + imageURL)
	if err != nil {
		return detectedFaces, fmt.Errorf("failed to read file: %s", err.Error())
	}

	if len(imageBytes) == 0 {
		return detectedFaces, fmt.Errorf("image data is empty")
	}

	resp, err := r.rekognitionClient.DetectFaces(&rekognition.DetectFacesInput{
		Image: &rekognition.Image{
			Bytes: imageBytes,
		},
		Attributes: []*string{
			aws.String(rekognition.AttributeAll),
		},
	})
	if err != nil {
		return detectedFaces, fmt.Errorf("failed to detect faces: %s", err.Error())
	}

	if len(resp.FaceDetails) == 0 {
		return detectedFaces, nil
	}

	for _, face := range resp.FaceDetails {
		detectedFaces = append(detectedFaces, port.DetectedFace{
			Confidence: *face.Confidence,
		})
	}

	return detectedFaces, nil
}
