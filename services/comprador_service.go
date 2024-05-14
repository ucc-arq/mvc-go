package services

import (
	compradorCliente "mvc-go/clients/comprador"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
)

type compradorService struct{}

type compradorServiceInterface interface {
	GetCompradorById(id int) (dto.CompradorDto, e.ApiError)
}

var (
	CompradorService compradorServiceInterface
)

func init() {
	CompradorService = &compradorService{}
}

func (s *compradorService) GetCompradorById(id int) (dto.CompradorDto, e.ApiError) {

	var comprador model.Comprador = compradorCliente.GetCompradorById(id)
	var compradorDto dto.CompradorDto

	compradorDto.Id = comprador.Id
	compradorDto.Nombre = comprador.Nombre

	return compradorDto, nil
}
