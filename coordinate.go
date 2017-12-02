package goutils

type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type RedisCoordinate struct {
	Coordinate
	Seconds int64 `json:"seconds"`
}
