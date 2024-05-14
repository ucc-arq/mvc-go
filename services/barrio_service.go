package services

import (
	barrioCliente "mvc-go/clients/barrio"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
)

type barrioService struct{}

type barrioServiceInterface interface {
	GetBarrioDetalleById(id int) (dto.BarrioDetailDto, e.ApiError)
	GetBarrios() (dto.BarriosDetailDto, e.ApiError)
}

var (
	BarrioService barrioServiceInterface
)

func init() {
	BarrioService = &barrioService{}
}

func (s *barrioService) GetBarrioDetalleById(id int) (dto.BarrioDetailDto, e.ApiError) {

	var barrio model.Barrio = barrioCliente.GetBarrioSensoresById(id)
	var barrioDetailDto dto.BarrioDetailDto

	if barrio.Id == 0 {
		return barrioDetailDto, e.NewBadRequestApiError("barrio not found")
	}

	barrioDetailDto.Id = barrio.Id
	barrioDetailDto.Descripcion = barrio.Descripcion

	for _, sensor := range barrio.Sensores {
		var dtoSensor dto.SensorMinDto

		dtoSensor.Serial = sensor.Serial
		dtoSensor.Id = sensor.Id
		dtoSensor.Activo = "NO"
		if sensor.Activo {
			dtoSensor.Activo = "SI"
		}
		barrioDetailDto.SensoresMinDto = append(barrioDetailDto.SensoresMinDto, dtoSensor)
	}

	return barrioDetailDto, nil
}

func (s *barrioService) GetBarrios() (dto.BarriosDetailDto, e.ApiError) {

	var barrios model.Barrios = barrioCliente.GetBarrios()
	var barriosDto dto.BarriosDetailDto

	for _, barrio := range barrios {
		var barrioDto dto.BarrioDetailDto

		barrioDto.Id = barrio.Id
		barrioDto.Descripcion = barrio.Descripcion

		barriosDto = append(barriosDto, barrioDto)
	}

	return barriosDto, nil
}
