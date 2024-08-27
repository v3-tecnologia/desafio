package aws

import (
	"context"
	"fmt"
	gateway "github.com/HaroldoFV/desafio/internal/gateway"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"log"
	"os"
)

type RekognitionFaceRecognizer struct {
	client       *rekognition.Client
	collectionID string
}

func NewRekognitionFaceRecognizer(client *rekognition.Client, collectionID string) gateway.FaceRecognizerInterface {
	return &RekognitionFaceRecognizer{
		client:       client,
		collectionID: collectionID,
	}
}

func (r *RekognitionFaceRecognizer) RecognizeFace(imagePath string) (bool, error) {
	imageBytes, err := os.ReadFile(imagePath)
	if err != nil {
		return false, fmt.Errorf("failed to read image file: %w", err)
	}

	log.Printf("Attempting to detect faces in image: %s", imagePath)

	detectInput := &rekognition.DetectFacesInput{
		Image: &types.Image{
			Bytes: imageBytes,
		},
	}

	detectResult, err := r.client.DetectFaces(context.TODO(), detectInput)
	if err != nil {
		log.Printf("Error detecting faces: %v", err)
	} else {
		log.Printf("DetectFaces result: %+v", detectResult)
	}

	if detectResult == nil || len(detectResult.FaceDetails) == 0 {
		log.Println("No faces detected by DetectFaces. Proceeding with SearchFacesByImage anyway.")
	} else {
		log.Printf("Detected %d faces in the image", len(detectResult.FaceDetails))
	}

	searchInput := &rekognition.SearchFacesByImageInput{
		CollectionId: aws.String(r.collectionID),
		Image: &types.Image{
			Bytes: imageBytes,
		},
		MaxFaces:           aws.Int32(1),
		FaceMatchThreshold: aws.Float32(70.0),
	}

	searchResult, searchErr := r.client.SearchFacesByImage(context.TODO(), searchInput)
	if searchErr != nil {
		log.Printf("Error searching faces by image: %v", searchErr)
	} else {
		log.Printf("SearchFacesByImage result: %+v", searchResult)
	}

	if searchResult != nil && len(searchResult.FaceMatches) > 0 {
		log.Printf("Face recognized with similarity: %f", *searchResult.FaceMatches[0].Similarity)
		return true, nil
	}

	log.Println("No matching faces found. Attempting to index the face.")

	indexInput := &rekognition.IndexFacesInput{
		CollectionId: aws.String(r.collectionID),
		Image: &types.Image{
			Bytes: imageBytes,
		},
		MaxFaces: aws.Int32(1),
		DetectionAttributes: []types.Attribute{
			types.AttributeDefault,
		},
	}

	indexResult, indexErr := r.client.IndexFaces(context.TODO(), indexInput)
	if indexErr != nil {
		log.Printf("Error indexing face: %v", indexErr)
		return false, fmt.Errorf("face indexing failed: %w", indexErr)
	}

	log.Printf("IndexFaces result: %+v", indexResult)

	if indexResult != nil && len(indexResult.FaceRecords) > 0 {
		log.Printf("Face indexed successfully. Face ID: %s", *indexResult.FaceRecords[0].Face.FaceId)
		return false, nil
	}

	return false, fmt.Errorf("face was not recognized and could not be indexed")
}
