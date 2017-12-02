package goutils

type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type RedisCoordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Seconds   int64   `json:"seconds"`
	Nanos     int32   `json:"nanos"`
}
