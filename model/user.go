package model

type User struct {
	Id       int    `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(350);not null"`
	LastName string `gorm:"type:varchar(250);not null"`
	UserName string `gorm:"type:varchar(150);not null;unique"`
	Password string `gorm:"type:varchar(150);not null"`

	Address   Address `gorm:"foreignkey:AddressId"`
	AddressId int

	Telephones Telephones `gorm:"foreignKey:UserId"`
}

type Users []User
