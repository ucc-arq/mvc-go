package dto

type UserDto struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

type UsersDto []UserDto
