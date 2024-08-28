package device

import (
	"desafio-backend/pkg/errors"
	"desafio-backend/util"
	"time"
)

type Device struct {
	util.BaseModel
	MacAddress string `json:"macAddress"`
}

func (main Main) FindByMacAddress(macAddress string) (*Device, errors.Error) {
	var nDevice Device

	rows, err := main.db.Set("gorm:auto_preload", true).Raw(queryDeviceByMacAddress, macAddress).Rows()
	if err != nil {
		return nil, errors.NewError("Find Device data error", err.Error()).
			WithOperations("FindByMacAddress.Raw")
	}

	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	if errScan := main.db.ScanRows(rows, &nDevice); err != nil {
		return nil, errors.NewError("Scan Device data error", errScan.Error()).
			WithOperations("FindByMacAddress.ScanRows")
	}

	return &nDevice, nil

}

func (main Main) SaveDevice(device Device) (Device, errors.Error) {
	entity := Device{}

	row := main.db.Exec(Insert(device.Timestamp.Format(time.RFC3339), device.MacAddress)).
		Raw(queryDeviceByMacAddress, device.MacAddress).
		Row()

	if row == nil {
		return Device{}, errors.NewError("Save error", "Insert Error").
			WithOperations("Save.Scan")
	}

	dbError := row.Scan(&entity.ID, &entity.MacAddress, &entity.Timestamp)

	if dbError != nil {
		return Device{}, errors.NewError("Save error", dbError.Error()).
			WithOperations("Save.Scan")
	}

	return entity, nil
}
