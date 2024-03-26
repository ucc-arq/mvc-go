package medicion

import (
	"mvc-go/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func RegistrarMedicion(medicion model.Medicion) model.Medicion {
	result := Db.Create(&medicion)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("Medicion Registrada: ", medicion.Id)
	return medicion
}
