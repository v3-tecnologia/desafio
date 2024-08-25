package repositories

import (
	"github.com/ThalesMonteir0/desafio/internal/domain"
	"github.com/ThalesMonteir0/desafio/internal/repositories/convert"
	"github.com/ThalesMonteir0/desafio/internal/repositories/entity"
	"github.com/ThalesMonteir0/desafio/pkg/err_rest"
)

func (t telemetryRepository) FindDeviceByMAC(macAddress string) (domain.DeviceDomain, *err_rest.ErrRest) {
	var device entity.Devices
	sql := `SELECT ID, mac_address FROM devices WHERE mac_address = $1`

	err := t.db.QueryRow(sql, macAddress).Scan(&device.ID, &device.MacAddress)
	if err != nil {
		return domain.DeviceDomain{}, err_rest.NewBadRequestErr("unable get Device from Mac_address")
	}

	return convert.ConvertDeviceEntityToDomain(device), nil
}
