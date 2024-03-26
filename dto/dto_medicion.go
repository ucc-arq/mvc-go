package dto

type MedicionDto struct {
	Fecha       string  `json:"fecha"`
	SensorId    int     `json:"sensor_id"`
	Presion     float64 `json:"presion"`
	Temperatura float64 `json:"temperatura"`
}

type MedicionesDto []MedicionDto
