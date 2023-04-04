package address

import (
	"mvc-go/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetAddressById(id int) model.Address {
	var address model.Address

	Db.Where("id = ?", id).First(&address)
	log.Debug("Address: ", address)

	return address
}

func InsertAddress(address model.Address) model.Address {
	result := Db.Create(&address)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("Address Created: ", address.Id)
	return address
}
