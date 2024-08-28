package device

import (
	"fmt"
)

const queryDeviceByMacAddress = `SELECT id, mac_address, created_at as "timestamp" from public.device where mac_address = ?`

func Insert(createdAt string, macAddress string) string {
	return fmt.Sprintf(`INSERT INTO device (created_at, mac_address ) VALUES ('%s', '%s');`, createdAt, macAddress)
}
