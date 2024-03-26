package app

import (
	barrioController "mvc-go/controllers/barrio"
	medicionController "mvc-go/controllers/medicion"
	sensorController "mvc-go/controllers/sensor"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	// Sensores Mapping
	router.POST("/sensor", sensorController.SensorInsert)
	router.PUT("/sensor/:id/activar", sensorController.ActivarSensor)
	router.PUT("/sensor/:id/pausar", sensorController.PausarSensor)
	router.GET("/sensor/:id", sensorController.GetSensorById)

	// Bario Mapping
	router.GET("/barrio/:id", barrioController.GetBarrioDetalleById)

	// Medidcion Mapping
	router.POST("/medicion", medicionController.RegistrarMedicion)

	log.Info("Finishing mappings configurations")
}
