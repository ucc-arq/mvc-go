package dto

type BarrioDetailDto struct {
	Id          int    `json:"id"`
	Descripcion string `json:"barrio"`

	SensoresMinDto SensoresMinDto `json:"sensores,omitempty"`
}

type BarriosDetailDto []BarrioDetailDto
