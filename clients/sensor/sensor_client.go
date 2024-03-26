package sensor

import (
	"mvc-go/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetSensorById(id int) model.Sensor {
	var sensor model.Sensor

	Db.Where("id = ?", id).Preload("Barrio").Preload("Mediciones").First(&sensor)
	log.Debug("Sensor: ", sensor)

	return sensor
}

func InsertSensor(sensor model.Sensor) model.Sensor {
	result := Db.Create(&sensor)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("Sensor Created: ", sensor.Id)
	return sensor
}

func ActivarSensor(id int) model.Sensor {
	var sensor model.Sensor
	Db.Where("id = ?", id).First(&sensor)
	log.Debug("Sensor: ", sensor)

	sensor.Activo = true
	result := Db.Save(&sensor)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("Sensor Activado: ", sensor.Id)
	return sensor
}

func PausarSensor(id int) model.Sensor {
	var sensor model.Sensor
	Db.Where("id = ?", id).First(&sensor)
	log.Debug("Sensor: ", sensor)

	sensor.Activo = false
	result := Db.Save(&sensor)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("Sensor Pausado: ", sensor.Id)
	return sensor
}
