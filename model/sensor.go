package model

// Domain Classes - "Sensor" entity
type Sensor struct {
	Id     int    `gorm:"primaryKey"`
	Serial string `gorm:"type:varchar(350);not null"`
	Activo bool   `gorm:"not null"`

	Barrio   Barrio `gorm:"foreignkey:BarrioId"`
	BarrioId int

	Mediciones Mediciones `gorm:"foreignKey:SensorId"`
}

type Sensores []Sensor
