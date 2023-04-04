package model

type Address struct {
	Id     int    `gorm:"primaryKey"`
	Street string `gorm:"type:varchar(350);not null"`
	Number int    `gorm:"not null"`
}

type Addresses []Address
