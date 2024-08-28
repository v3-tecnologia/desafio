package gps

import "fmt"

const queryGpsById = `SELECT d.mac_address,
       						 g.timestamp,
       						 ST_AsGeoJSON(g.coordinates)
					    from public.device d, 
					         public.gps g
					   WHERE g.id = ? 
					     and g.device_id = d.id`

func Insert(deviceId int, timestamp string, latitude, longitude float64) string {
	return fmt.Sprintf("INSERT INTO gps  (device_id, timestamp, coordinates ) VALUES ('%d', '%s', ST_GeogFromText('SRID=4326;POINT(%f %f)')) RETURNING id;", deviceId, timestamp, latitude, longitude)
}
