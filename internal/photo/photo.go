package photo

import (
	"desafio-backend/pkg/errors"
	"desafio-backend/pkg/logger"
	"desafio-backend/util"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var (
	_, b, _, _     = runtime.Caller(0)
	RootPath       = filepath.Join(filepath.Dir(b), "../../resources")
	tempFolderPath = filepath.Join(RootPath, "/temp-files")
)

type ImageFile struct {
	Name     string
	FullPath string
	MimeType string
	Bytes    []byte
}

type Request struct {
	MacAddress string    `json:"macAddress"`
	Timestamp  time.Time `json:"timestamp"`
	ImageFile  ImageFile `json:"ImageFile"`
}

type Response struct {
	MacAddress string    `json:"macAddress"`
	Timestamp  time.Time `json:"timestamp"`
	ImageFile  ImageFile `json:"latitude"`
}

func (main Main) ParseImage(file multipart.File, fileHeader *multipart.FileHeader) (ImageFile, errors.Error) {
	defer file.Close()

	tempFileName := fmt.Sprintf("uploaded-%s-*%s", removedExt(fileHeader.Filename), filepath.Ext(fileHeader.Filename))

	tempFile, err := os.CreateTemp(tempFolderPath, tempFileName)
	if err != nil {
		logger.Error(util.GeneralParseError, "ParseImage", err, fileHeader)
		return ImageFile{}, errors.NewError("Error in creating the file", err.Error()).
			WithOperations("ParseImage.CreateTemp")
	}

	defer tempFile.Close()

	filebytes, err := io.ReadAll(file)
	if err != nil {
		logger.Error(util.GeneralParseError, "ParseImage", err, fileHeader)
		return ImageFile{}, errors.NewError("Error in reading the file buffer", err.Error()).
			WithOperations("ParseImage.ReadAll")
	}

	tempFile.Write(filebytes)

	_, tFilename := filepath.Split(tempFile.Name())

	imgFile := ImageFile{
		Name:     tFilename,
		FullPath: tempFile.Name(),
		MimeType: fileHeader.Header.Get("Content-Type"),
		Bytes:    filebytes,
	}

	return imgFile, nil
}

func (main Main) ParsePhoto(photo string, file ImageFile) (Request, errors.Error) {
	var nPhoto *Request

	err := json.Unmarshal([]byte(photo), &nPhoto)

	nPhoto.ImageFile = file

	if err != nil {
		logger.Error(util.GeneralParseError, "ParsePhoto", err, photo)
		return Request{}, errors.NewError("Connot decode data", err.Error()).
			WithOperations("ParsePhoto.Decode")
	}
	return *nPhoto, nil
}

func (main Main) SavePhoto(photo Request) (Response, errors.Error) {
	// TODO save the data received into a database
	return photo.toResponse(), nil
}

func (main Main) ValidatePhoto(photo Request) errors.ErrorList {
	// TODO validate data received
	return nil
}

func (entity Request) toResponse() Response {
	return Response{
		MacAddress: entity.MacAddress,
		Timestamp:  entity.Timestamp,
		ImageFile:  entity.ImageFile,
	}
}

func removedExt(file string) string {
	return strings.TrimSuffix(file, filepath.Ext(file))
}
