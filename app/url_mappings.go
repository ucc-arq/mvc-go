package app

import (
	barrioController "mvc-go/controllers/barrio"
	medicionController "mvc-go/controllers/medicion"
	sensorController "mvc-go/controllers/sensor"
	vendedorController "mvc-go/controllers/vendedor"

	compradorController "mvc-go/controllers/comprador"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	// Sensores Mapping
	router.POST("/sensor", sensorController.SensorInsert)
	router.PUT("/sensor/:id/activar", sensorController.ActivarSensor)
	router.PUT("/sensor/:id/pausar", sensorController.PausarSensor)
	router.GET("/sensor/:id", sensorController.GetSensorById)

	router.GET("/vendedor/:id", vendedorController.GetVendedorById)

	router.GET("/comprador/:id", compradorController.GetCompradorById)

	// Bario Mapping
	router.GET("/barrio/:id", barrioController.GetBarrioDetalleById)
	router.GET("/barrio", barrioController.GetBarrios)

	// Medidcion Mapping
	router.POST("/medicion", medicionController.RegistrarMedicion)

	log.Info("Finishing mappings configurations")
}
