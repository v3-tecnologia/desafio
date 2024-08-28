package photo

import (
	"desafio-backend/internal/device"
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
	var ID int64

	processedDevice, deviceErr := main.processAndSaveDevice(photo.MacAddress)
	if deviceErr != nil {
		return Response{}, deviceErr
	}

	rows, err := main.db.Raw(Insert(processedDevice.ID, photo.Timestamp.Format(time.RFC3339), photo.ImageFile.FullPath, photo.ImageFile.Name)).Rows()

	if err != nil {
		return Response{}, errors.NewError("Save photo error", err.Error()).
			WithOperations("SavePhoto.Raw")
	}

	defer rows.Close()

	if errScan := main.db.ScanRows(rows, &ID); err != nil {
		return Response{}, errors.NewError("Scan photo data error", errScan.Error()).
			WithOperations("SavePhoto.ScanRows")
	}

	return main.findGyroscopeById(ID)
}

func (main Main) ValidatePhoto(photo Request) errors.ErrorList {
	ers := errors.NewErrorList()

	if photo.MacAddress == "" {
		err := errors.NewError("Missing MacAddress", "MacAddress is required").
			WithMeta("field", "macAddress").
			WithOperations("ValidatePhoto.CheckMacAddress")
		ers.Append(err)
	}

	macErr := util.IsValidateMacAddress(photo.MacAddress)
	if macErr != nil {
		ers.Append(errors.NewError("Invalid MacAddress format", "MacAddress is not valid").
			WithMeta("field", "macAddress").
			WithOperations("ValidatePhoto.MacAddressFormat"))
	}

	if photo.Timestamp.IsZero() {
		err := errors.NewError("Missing Timestamp", "Timestamp is required").
			WithMeta("field", "timestamp").
			WithOperations("ValidateGps.CheckTimestamp")
		ers.Append(err)
	}

	return ers
}

func removedExt(file string) string {
	return strings.TrimSuffix(file, filepath.Ext(file))
}

func (main Main) processAndSaveDevice(macAddress string) (*device.Device, errors.Error) {
	// try to find a foundDevice with the macAddress
	foundDevice, deviceErr := main.deviceMain.FindByMacAddress(macAddress)

	if deviceErr != nil {
		return &device.Device{}, deviceErr
	}

	// if a device is not found, then insert it
	if foundDevice == nil {
		var insertDevice = device.Device{}

		insertDevice.Timestamp = time.Now()
		insertDevice.MacAddress = macAddress
		insertDevice, err := main.deviceMain.SaveDevice(insertDevice)

		if err != nil {
			return &device.Device{}, err
		}

		return &insertDevice, nil
	}

	return foundDevice, nil
}

func (main Main) findGyroscopeById(ID int64) (Response, errors.Error) {
	response := Response{}
	row := main.db.Set("gorm:auto_preload", true).Raw(queryGyroscopeById, ID).Row()

	if errScan := row.Scan(&response.MacAddress, &response.Timestamp, &response.ImageFile.FullPath, &response.ImageFile.Name); errScan != nil {
		return Response{}, errors.NewError("Scan GPS data error", errScan.Error()).
			WithOperations("SaveGps.ScanRows")
	}

	return response, nil
}
