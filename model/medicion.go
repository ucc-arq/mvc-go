package model

// Domain Classes - "Sensor" entity
type Medicion struct {
	Id          int     `gorm:"primaryKey"`
	Fecha       string  `gorm:"type:varchar(250);not null"`
	Presion     float64 `gorm:"not null"`
	Temperatura float64 `gorm:"not null"`

	Sensor   Sensor `gorm:"foreignkey:SensorId"`
	SensorId int
}

type Mediciones []Medicion
