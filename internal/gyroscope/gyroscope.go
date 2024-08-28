package gyroscope

import (
	"desafio-backend/internal/device"
	"desafio-backend/pkg/errors"
	"desafio-backend/pkg/logger"
	"desafio-backend/util"
	"encoding/json"
	"io"
	"time"
)

// Gyroscope limit in  (°/s) - ±2000°/s is an example value for high sensibility gyroscopes
const (
	minValue = -2000.0
	maxValue = 2000.0
)

type Request struct {
	MacAddress string    `json:"macAddress"`
	Timestamp  time.Time `json:"timestamp"`
	XAxis      float64   `json:"xAxis"`
	YAxis      float64   `json:"yAxis"`
	ZAxis      float64   `json:"zAxis"`
}

type Response struct {
	MacAddress string    `json:"macAddress"`
	Timestamp  time.Time `json:"timestamp"`
	XAxis      float64   `json:"xAxis"`
	YAxis      float64   `json:"yAxis"`
	ZAxis      float64   `json:"zAxis"`
}

func (main Main) ParseGyroscope(gyroscope io.ReadCloser) (Request, errors.Error) {
	var nGyroscope *Request

	err := json.NewDecoder(gyroscope).Decode(&nGyroscope)
	if err != nil {
		logger.Error(util.GeneralParseError, "ParseGyroscope", err, gyroscope)
		return Request{}, errors.NewError("Connot decode data", err.Error()).
			WithOperations("ParseGyroscope.Decode")
	}
	return *nGyroscope, nil
}

func (main Main) SaveGyroscope(gyroscope Request) (Response, errors.Error) {
	var ID int64

	processedDevice, deviceErr := main.processAndSaveDevice(gyroscope.MacAddress)
	if deviceErr != nil {
		return Response{}, deviceErr
	}

	rows, err := main.db.Raw(Insert(processedDevice.ID, gyroscope.Timestamp.Format(time.RFC3339), gyroscope.XAxis, gyroscope.YAxis, gyroscope.ZAxis)).Rows()

	if err != nil {
		return Response{}, errors.NewError("Save Gyroscope error", err.Error()).
			WithOperations("SaveGyroscope.Raw")
	}

	defer rows.Close()

	if errScan := main.db.ScanRows(rows, &ID); err != nil {
		return Response{}, errors.NewError("Scan Gyroscope data error", errScan.Error()).
			WithOperations("SaveGyroscope.ScanRows")
	}

	return main.findGyroscopeById(ID)
}

func (main Main) ValidateGyroscope(gyroscope Request) errors.ErrorList {
	ers := errors.NewErrorList()

	if gyroscope.MacAddress == "" {
		ers.Append(errors.NewError("Missing MacAddress", "MacAddress is required").
			WithMeta("field", "macAddress").
			WithOperations("ValidateGyroscope.MacAddress"))
	}

	macErr := util.IsValidateMacAddress(gyroscope.MacAddress)
	if macErr != nil {
		ers.Append(errors.NewError("Invalid MacAddress format", "MacAddress is not valid").
			WithMeta("field", "macAddress").
			WithOperations("ValidateGyroscope.MacAddressFormat"))
	}

	if gyroscope.Timestamp.IsZero() {
		ers.Append(errors.NewError("Missing Timestamp", "Timestamp is required").
			WithMeta("field", "timestamp").
			WithOperations("ValidateGyroscope.Timestamp"))
	}

	if gyroscope.XAxis == 0 && gyroscope.YAxis == 0 && gyroscope.ZAxis == 0 {
		ers.Append(errors.NewError("Invalid Gyroscope Data", "At least one axis value must be non-zero").
			WithOperations("ValidateGyroscope.GyroscopeData"))
	}

	if gyroscope.XAxis < minValue || gyroscope.XAxis > maxValue {
		ers.Append(errors.NewError("Invalid X axis", "X axis value must be within defined range").
			WithOperations("ValidateGyroscope.XAxis"))
	}

	if gyroscope.YAxis < minValue || gyroscope.YAxis > maxValue {
		ers.Append(errors.NewError("Invalid Y axis", "Y axis value must be within defined range").
			WithOperations("ValidateGyroscope.YAxis"))
	}

	if gyroscope.ZAxis < minValue || gyroscope.ZAxis > maxValue {
		ers.Append(errors.NewError("Invalid Z axis", "Z axis value must be within defined range").
			WithOperations("ValidateGyroscope.ZAxis"))
	}

	return ers
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

	if errScan := row.Scan(&response.MacAddress, &response.Timestamp, &response.XAxis, &response.YAxis, &response.ZAxis); errScan != nil {
		return Response{}, errors.NewError("Scan GPS data error", errScan.Error()).
			WithOperations("SaveGps.ScanRows")
	}

	return response, nil
}
