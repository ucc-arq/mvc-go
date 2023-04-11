package telephone

import (
	"mvc-go/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func AddTelephone(telephone model.Telephone) model.Telephone {
	result := Db.Create(&telephone)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("Telephone Created: ", telephone.Id)
	return telephone
}
