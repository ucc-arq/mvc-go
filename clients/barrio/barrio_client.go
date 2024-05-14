package barrio

import (
	"mvc-go/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetBarrioById(id int) model.Barrio {
	var barrio model.Barrio

	Db.Where("id = ?", id).First(&barrio)
	log.Debug("Barrio: ", barrio)

	return barrio
}

func GetBarrioSensoresById(id int) model.Barrio {
	var barrio model.Barrio

	Db.Where("id = ?", id).Preload("Sensores").First(&barrio)
	log.Debug("Barrio: ", barrio)

	return barrio
}

func GetBarrios() model.Barrios {
	var barrios model.Barrios
	Db.Order("descripcion").Find(&barrios)

	log.Debug("Barrios: ", barrios)

	return barrios
}
