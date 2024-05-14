package vendedor

import (
	"mvc-go/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetVendedorById(id int) model.Vendedor {
	var vendedor model.Vendedor

	Db.Where("id = ?", id).First(&vendedor)
	log.Debug("Vendedor: ", vendedor)

	return vendedor
}
