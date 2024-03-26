package services

import (
	barrioCliente "mvc-go/clients/barrio"
	sensorCliente "mvc-go/clients/sensor"

	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
)

type sensorService struct{}

type sensorServiceInterface interface {
	GetSensorById(id int) (dto.SensorDetailDto, e.ApiError)
	InsertSensor(sensorDto dto.SensorDto) (dto.SensorDto, e.ApiError)
	ActivarSensor(id int) (dto.SensorDto, e.ApiError)
	PausarSensor(id int) (dto.SensorDto, e.ApiError)
}

var (
	SensorService sensorServiceInterface
)

func init() {
	SensorService = &sensorService{}
}

func (s *sensorService) GetSensorById(id int) (dto.SensorDetailDto, e.ApiError) {

	var sensor model.Sensor = sensorCliente.GetSensorById(id)
	var sensorDetailDto dto.SensorDetailDto

	if sensor.Id == 0 {
		return sensorDetailDto, e.NewBadRequestApiError("sensor not found")
	}

	sensorDetailDto.Serial = sensor.Serial
	sensorDetailDto.Id = sensor.Id
	sensorDetailDto.Barrio = sensor.Barrio.Descripcion

	sensorDetailDto.Activo = "NO"
	if sensor.Activo {
		sensorDetailDto.Activo = "SI"
	}

	for _, medicion := range sensor.Mediciones {
		var dtoMedicion dto.MedicionMinDto

		dtoMedicion.Fecha = medicion.Fecha
		dtoMedicion.Presion = medicion.Presion
		dtoMedicion.Temperatura = medicion.Temperatura

		sensorDetailDto.MedicionesMinDto = append(sensorDetailDto.MedicionesMinDto, dtoMedicion)
	}

	return sensorDetailDto, nil
}

func (s *sensorService) InsertSensor(sensorDto dto.SensorDto) (dto.SensorDto, e.ApiError) {

	var sensor model.Sensor
	var barrio model.Barrio
	sensor.Serial = sensorDto.Serial

	barrio = barrioCliente.GetBarrioById(sensorDto.BarrioId)

	sensor.Barrio = barrio
	sensor.BarrioId = barrio.Id
	sensor.Activo = true

	sensor = sensorCliente.InsertSensor(sensor)

	sensorDto.Id = sensor.Id

	return sensorDto, nil
}

func (s *sensorService) ActivarSensor(id int) (dto.SensorDto, e.ApiError) {
	var sensor model.Sensor
	var sensorDto dto.SensorDto

	sensor = sensorCliente.ActivarSensor(id)

	sensorDto.Id = sensor.Id
	sensorDto.Activo = "SI"

	return sensorDto, nil
}

func (s *sensorService) PausarSensor(id int) (dto.SensorDto, e.ApiError) {
	var sensor model.Sensor
	var sensorDto dto.SensorDto

	sensor = sensorCliente.PausarSensor(id)

	sensorDto.Id = sensor.Id
	sensorDto.Activo = "NO"

	return sensorDto, nil
}
