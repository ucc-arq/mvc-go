package db

import (
	barrioClient "mvc-go/clients/barrio"
	compradorClient "mvc-go/clients/comprador"
	medicionClient "mvc-go/clients/medicion"
	sensorClient "mvc-go/clients/sensor"
	vendedorClient "mvc-go/clients/vendedor"

	"mvc-go/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// DB Connections Paramters
	DBName := "uccarqsoft"  //Nombre de la base de datos local de ustedes
	DBUser := "egaite"      //usuario de la base de datos, habitualmente root
	DBPass := ""            //password del root en la instalacion
	DBHost := "db4free.net" //host de la base de datos. hbitualmente 127.0.0.1
	// ------------------------

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	// We need to add all CLients that we build
	barrioClient.Db = db
	sensorClient.Db = db
	medicionClient.Db = db
	vendedorClient.Db = db
	compradorClient.Db = db
}

func StartDbEngine() {
	// We need to migrate all classes model.
	db.AutoMigrate(&model.Barrio{})
	db.AutoMigrate(&model.Sensor{})
	db.AutoMigrate(&model.Medicion{})

	db.AutoMigrate(&model.Vendedor{})
	db.AutoMigrate(&model.Comprador{})

	log.Info("Finishing Migration Database Tables")
}
