package usecase

import (
	"errors"
	"fmt"
	"github.com/HaroldoFV/desafio/configs"
	domain "github.com/HaroldoFV/desafio/internal/domain"
	entity "github.com/HaroldoFV/desafio/internal/domain/entity"
	"github.com/HaroldoFV/desafio/internal/dto"
	"github.com/HaroldoFV/desafio/internal/gateway"
	"github.com/google/uuid"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type CreatePhotoUseCase struct {
	PhotoRepository domain.PhotoRepositoryInterface
	FaceRecognizer  gateway.FaceRecognizerInterface
	Config          *configs.Conf
}

func NewCreatePhotoUseCase(
	photoRepository domain.PhotoRepositoryInterface,
	faceRecognizer gateway.FaceRecognizerInterface,
	config *configs.Conf,
) *CreatePhotoUseCase {
	return &CreatePhotoUseCase{
		PhotoRepository: photoRepository,
		FaceRecognizer:  faceRecognizer,
		Config:          config,
	}
}

func (c *CreatePhotoUseCase) Execute(input dto.CreatePhotoInputDTO) (dto.PhotoOutputDTO, error) {
	if !isValidImageType(input.Image.Filename) {
		return dto.PhotoOutputDTO{}, errors.New("invalid image type")
	}

	if input.FileSize > 5<<20 { // 5 MB
		return dto.PhotoOutputDTO{}, fmt.Errorf("file is too large. Maximum size is 5MB")
	}

	fileName := generateUniqueFileName(input.Image.Filename)
	filePath := filepath.Join(c.Config.PhotoStoragePath, fileName)

	err := saveFile(input.Image, filePath)
	if err != nil {
		return dto.PhotoOutputDTO{}, err
	}

	// Perform face recognition
	recognized, err := c.FaceRecognizer.RecognizeFace(filePath)
	if err != nil {
		log.Printf("Error in face recognition: %v", err)
		recognized = false
	}

	photo, err := entity.NewPhoto(filePath, input.MacAddress)
	if err != nil {
		removeErr := os.Remove(filePath)
		if removeErr != nil {
			log.Printf("error removing file: %v", removeErr)
		}
		return dto.PhotoOutputDTO{}, err
	}

	if err = c.PhotoRepository.Create(photo); err != nil {
		removeErr := os.Remove(filePath)
		if removeErr != nil {
			log.Printf("error removing file: %v", removeErr)
		}
		return dto.PhotoOutputDTO{}, err
	}
	outputDTO := dto.PhotoOutputDTO{
		ID:         photo.GetID(),
		Timestamp:  photo.GetTimestamp(),
		MacAddress: photo.GetMACAddress(),
		FilePath:   photo.GetFilePath(),
		Recognized: recognized,
	}
	return outputDTO, nil
}

func isValidImageType(filename string) bool {
	ext := filepath.Ext(filename)
	validExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
	}
	return validExtensions[ext]
}

func generateUniqueFileName(originalFilename string) string {
	ext := filepath.Ext(originalFilename)
	nameWithoutExt := strings.TrimSuffix(originalFilename, ext)
	timestamp := time.Now().Format("20060102_150405")
	uniqueID := uuid.New().String()
	uniqueFileName := fmt.Sprintf("%s_%s_%s%s", nameWithoutExt, timestamp, uniqueID, ext)
	uniqueFileName = strings.ReplaceAll(uniqueFileName, " ", "_")

	return uniqueFileName
}

func saveFile(file *multipart.FileHeader, destPath string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	return err
}
