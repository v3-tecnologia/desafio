package gyroscope

import "fmt"

const queryGyroscopeById = `SELECT d.mac_address,
       						 g.timestamp,
       						 g.x_axis,
       						 g.y_axis,
       						 g.z_axis
					    from public.device d, 
					         public.gyroscope g
					   WHERE g.id = ? 
					     and g.device_id = d.id`

func Insert(deviceId int, timestamp string, xAxis, yAxis, zAxis float64) string {
	return fmt.Sprintf("INSERT INTO gyroscope  (device_id, timestamp, x_axis, y_axis, z_axis ) VALUES ('%d', '%s', '%f',  '%f', '%f') RETURNING id;", deviceId, timestamp, xAxis, yAxis, zAxis)
}
