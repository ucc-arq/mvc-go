package dto

type VendedorDto struct {
	Id       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Dni      string `json:"dnbi"`
}

type VendeodresDto []VendedorDto
