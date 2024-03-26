package services

import (
	medicionCliente "mvc-go/clients/medicion"
	sensorCliente "mvc-go/clients/sensor"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
	"time"
)

type medicionService struct{}

type medicionServiceInterface interface {
	RegistrarMedicion(medicionDto dto.MedicionDto) (dto.MedicionDto, e.ApiError)
}

var (
	MedicionService medicionServiceInterface
)

func init() {
	MedicionService = &medicionService{}
}

func (s *medicionService) RegistrarMedicion(medicionDto dto.MedicionDto) (dto.MedicionDto, e.ApiError) {

	var medicion model.Medicion

	var sensor model.Sensor = sensorCliente.GetSensorById(medicionDto.SensorId)

	if !sensor.Activo {
		return medicionDto, e.NewBadRequestApiError("Sensor Inactivo")
	}

	medicion.Presion = medicionDto.Presion
	medicion.SensorId = medicionDto.SensorId
	medicion.Temperatura = medicionDto.Temperatura

	now := time.Now()
	formatted := now.Format(time.RFC822)

	medicion.Fecha = formatted
	medicion = medicionCliente.RegistrarMedicion(medicion)

	//medicion.Id = medicion.Id

	return medicionDto, nil
}
