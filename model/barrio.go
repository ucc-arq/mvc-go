package model

// Domain Classes - "Barrio" entity
type Barrio struct {
	Id          int    `gorm:"primaryKey"`
	Descripcion string `gorm:"type:varchar(350);not null"`

	Sensores Sensores `gorm:"foreignKey:BarrioId"`
}

type Barrios []Barrio
