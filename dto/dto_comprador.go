package dto

type CompradorDto struct {
	Id       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Dni      string `json:"dni"`
}

type CompradoresDto []CompradorDto
