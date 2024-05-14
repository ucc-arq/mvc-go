package services

import (
	vendedorCliente "mvc-go/clients/vendedor"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
)

type vendedorService struct{}

type vendedorServiceInterface interface {
	GetVendedorById(id int) (dto.VendedorDto, e.ApiError)
}

var (
	VendedorService vendedorServiceInterface
)

func init() {
	VendedorService = &vendedorService{}
}

func (s *vendedorService) GetVendedorById(id int) (dto.VendedorDto, e.ApiError) {

	var vendedor model.Vendedor = vendedorCliente.GetVendedorById(id)
	var vendedorDto dto.VendedorDto

	vendedorDto.Id = vendedor.Id
	vendedorDto.Nombre = vendedor.Nombre

	return vendedorDto, nil
}
