package model

// Domain Classes - "Barrio" entity
type Comprador struct {
	Id       int    `gorm:"primaryKey"`
	Nombre   string `gorm:"type:varchar(600);not null"`
	Apellido string `gorm:"type:varchar(350);not null"`
	Dni      string `gorm:"type:varchar(12);not null"`
}

type Compradores []Comprador
