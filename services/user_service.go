package services

import (
	addressCliente "mvc-go/clients/address"
	userCliente "mvc-go/clients/user"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
)

type userService struct{}

type userServiceInterface interface {
	GetUserById(id int) (dto.UserDto, e.ApiError)
	GetUsers() (dto.UsersDto, e.ApiError)
	InsertUser(userDto dto.UserDto) (dto.UserDto, e.ApiError)
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

func (s *userService) GetUserById(id int) (dto.UserDto, e.ApiError) {

	var user model.User = userCliente.GetUserById(id)
	var userDto dto.UserDto

	if user.Id == 0 {
		return userDto, e.NewBadRequestApiError("user not found")
	}
	userDto.Name = user.Name
	userDto.LastName = user.LastName
	userDto.UserName = user.UserName
	userDto.Id = user.Id
	userDto.Street = user.Address.Street
	userDto.Number = user.Address.Number

	return userDto, nil
}

func (s *userService) GetUsers() (dto.UsersDto, e.ApiError) {

	var users model.Users = userCliente.GetUsers()
	var usersDto dto.UsersDto

	for _, user := range users {
		var userDto dto.UserDto
		userDto.Name = user.Name
		userDto.LastName = user.LastName
		userDto.UserName = user.Name
		userDto.Id = user.Id

		userDto.Street = user.Address.Street
		userDto.Number = user.Address.Number

		usersDto = append(usersDto, userDto)
	}

	return usersDto, nil
}

func (s *userService) InsertUser(userDto dto.UserDto) (dto.UserDto, e.ApiError) {

	var user model.User

	var address model.Address

	user.Name = userDto.Name
	user.LastName = userDto.LastName
	user.UserName = userDto.UserName
	user.Password = userDto.Password

	address.Number = userDto.Number
	address.Street = userDto.Street
	address = addressCliente.InsertAddress(address)

	user.Address = address
	user = userCliente.InsertUser(user)

	userDto.Id = user.Id

	return userDto, nil
}
