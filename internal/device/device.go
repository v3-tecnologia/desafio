package device

import (
	"desafio-backend/pkg/errors"
	"desafio-backend/util"
)

type Device struct {
	util.BaseModel
	MacAddress string `json:"macAddress"`
}

func (main Main) FindByMacAddress(macAddress string) (Device, errors.Error) {
	var device Device

	rows, err := main.db.Set("gorm:auto_preload", true).Raw(queryDeviceByMacAddress, macAddress).Rows()
	if err != nil {
		return Device{}, errors.NewError("Find Device data error", err.Error()).
			WithOperations("FindByMacAddress.Raw")
	}

	defer rows.Close()

	if !rows.Next() {
		return Device{}, errors.NewError("Device not found", "Device with mac address "+macAddress+" not found").
			WithOperations("FindByMacAddress.Rows.Next")
	}

	if errScan := main.db.ScanRows(rows, &device); err != nil {
		return Device{}, errors.NewError("Scan Device data error", errScan.Error()).
			WithOperations("FindByMacAddress.ScanRows")
	}

	return device, nil

}

func (main Main) SaveDevice(device Device) (Device, errors.Error) {
	if err := main.db.Create(&device).Error; err != nil {
		return Device{}, errors.NewError("Save Device data error", err.Error()).
			WithOperations("SaveDevice.Create")
	}
	return device, nil
}
