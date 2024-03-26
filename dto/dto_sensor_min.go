package dto

type SensorMinDto struct {
	Serial string `json:"serial"`
	Id     int    `json:"id"`
	Activo string `json:"activo"`
}

type SensoresMinDto []SensorMinDto
