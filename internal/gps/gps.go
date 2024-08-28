package gps

import (
	"desafio-backend/internal/device"
	"desafio-backend/pkg/errors"
	"desafio-backend/pkg/logger"
	"desafio-backend/util"
	"encoding/json"
	"io"
	"time"
)

type Request struct {
	MacAddress string    `json:"macAddress"`
	Timestamp  time.Time `json:"timestamp"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
}

type Response struct {
	MacAddress  string    `json:"macAddress"`
	Timestamp   time.Time `json:"timestamp"`
	Coordinates string    `json:"coordinates"`
}

func (main Main) ParseGps(gps io.ReadCloser) (Request, errors.Error) {
	var nGps *Request

	err := json.NewDecoder(gps).Decode(&nGps)
	if err != nil {
		logger.Error(util.GeneralParseError, "ParseGps", err, gps)
		return Request{}, errors.NewError("Connot decode data", err.Error()).
			WithOperations("ParseGps.Decode")
	}
	return *nGps, nil
}

func (main Main) SaveGps(gps Request) (Response, errors.Error) {
	var ID int64

	processedDevice, deviceErr := main.processAndSaveDevice(gps.MacAddress)
	if deviceErr != nil {
		return Response{}, deviceErr
	}

	rows, err := main.db.Raw(Insert(processedDevice.ID, gps.Timestamp.Format(time.RFC3339), gps.Latitude, gps.Longitude)).Rows()

	if err != nil {
		return Response{}, errors.NewError("Save GPS error", err.Error()).
			WithOperations("SaveGps.Raw")
	}

	defer rows.Close()

	if errScan := main.db.ScanRows(rows, &ID); err != nil {
		return Response{}, errors.NewError("Scan GPS data error", errScan.Error()).
			WithOperations("SaveGps.ScanRows")
	}

	return main.findGpsById(ID)
}

func (main Main) findGpsById(ID int64) (Response, errors.Error) {
	response := Response{}
	row := main.db.Set("gorm:auto_preload", true).Raw(queryGpsById, ID).Row()

	if errScan := row.Scan(&response.MacAddress, &response.Timestamp, &response.Coordinates); errScan != nil {
		return Response{}, errors.NewError("Scan GPS data error", errScan.Error()).
			WithOperations("SaveGps.ScanRows")
	}

	return response, nil
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

func (main Main) ValidateGps(gps Request) errors.ErrorList {
	ers := errors.NewErrorList()

	if gps.MacAddress == "" {
		err := errors.NewError("Missing MacAddress", "MacAddress is required").
			WithMeta("field", "macAddress").
			WithOperations("ValidateGps.CheckMacAddress")
		ers.Append(err)
	}

	macErr := util.IsValidateMacAddress(gps.MacAddress)
	if macErr != nil {
		ers.Append(errors.NewError("Invalid MacAddress format", "MacAddress is not valid").
			WithMeta("field", "macAddress").
			WithOperations("ValidateGps.MacAddressFormat"))
	}

	if gps.Latitude == 0.0 {
		err := errors.NewError("Missing Latitude", "Latitude is required").
			WithMeta("field", "latitude").
			WithOperations("ValidateGps.CheckLatitude")
		ers.Append(err)
	}

	if gps.Longitude == 0.0 {
		err := errors.NewError("Missing Longitude", "Longitude is required").
			WithMeta("field", "longitude").
			WithOperations("ValidateGps.CheckLongitude")
		ers.Append(err)
	}

	if gps.Timestamp.IsZero() {
		err := errors.NewError("Missing Timestamp", "Timestamp is required").
			WithMeta("field", "timestamp").
			WithOperations("ValidateGps.CheckTimestamp")
		ers.Append(err)
	}

	return ers
}
