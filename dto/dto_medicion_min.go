package dto

type MedicionMinDto struct {
	Fecha       string  `json:"fecha"`
	Presion     float64 `json:"presion"`
	Temperatura float64 `json:"temperatura"`
}

type MedicionesMinDto []MedicionMinDto
