package device

const queryDeviceByMacAddress = `SELECT id, mac_address, created_at as "timestamp" from public.device where mac_address = ?`
