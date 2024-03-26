package dto

type SensorDetailDto struct {
	Serial string `json:"serial"`
	Id     int    `json:"id"`
	Barrio string `json:"barrio"`
	Activo string `json:"activo"`

	MedicionesMinDto MedicionesMinDto `json:"mediciones,omitempty"`
}

type SensoresDetailDto []SensorDetailDto
