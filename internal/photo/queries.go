package photo

import "fmt"

const queryGyroscopeById = `SELECT d.mac_address,
       						 p.timestamp,
       						 p.path,
       						 p.name
					    from public.device d, 
					         public.photo p
					   WHERE p.id = ? 
					     and p.device_id = d.id`

func Insert(deviceId int, timestamp string, path string, name string) string {
	return fmt.Sprintf("INSERT INTO photo (device_id, timestamp, path, name ) VALUES ('%d', '%s', '%s',  '%s') RETURNING id;", deviceId, timestamp, path, name)
}
