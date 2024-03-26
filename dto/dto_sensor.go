package dto

type SensorDto struct {
	Serial   string `json:"serial"`
	Id       int    `json:"id"`
	BarrioId int    `json:"barrio_id"`
	Activo   string `json:"activo"`
}

type SensoresDto []SensorDto
