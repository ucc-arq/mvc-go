package model

type Telephone struct {
	Id     int    `gorm:"primaryKey"`
	Code   string `gorm:"type:varchar(10);not null"`
	Number string `gorm:"type:varchar(25);not null"`

	UserId int
}

type Telephones []Telephone
