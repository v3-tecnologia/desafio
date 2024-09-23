package services

import (
    "BackendGo/config"
    "context"
    "fmt"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/rekognition"
    "log"
    "os"
)

func ComparePhotoWithRekognition(photoPath string) (bool, error) {
    // Configurar sessão AWS
    sess, err := session.NewSession(&aws.Config{
        Region:      aws.String(config.Config.AWSRegion),
        Credentials: credentials.NewStaticCredentials(config.Config.AWSAccessKey, config.Config.AWSSecretKey, ""),
    })
    if err != nil {
        return false, err
    }

    svc := rekognition.New(sess)

    file, err := os.Open(photoPath)
    if err != nil {
        return false, err
    }
    defer file.Close()

    // Carregar imagem
    fileInfo, _ := file.Stat()
    size := fileInfo.Size()
    buffer := make([]byte, size)
    file.Read(buffer)

    image := rekognition.Image{
        Bytes: buffer,
    }

    input := &rekognition.CompareFacesInput{
        SourceImage: &image,
        TargetImage: &image,
        SimilarityThreshold: aws.Float64(90.0),
    }

    result, err := svc.CompareFaces(input)
    if err != nil {
        return false, err
    }

    for _, face := range result.FaceMatches {
        fmt.Printf("Similarity: %f%%\n", *face.Similarity)
        if *face.Similarity > 90.0 {
            return true, nil
        }
    }

    return false, nil
}
