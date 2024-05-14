package comprador

import (
	"mvc-go/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetCompradorById(id int) model.Comprador {
	var comprador model.Comprador

	Db.Where("id = ?", id).First(&comprador)
	log.Debug("Comprador: ", comprador)

	return comprador
}
