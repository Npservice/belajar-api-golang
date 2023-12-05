package service

import (
	"perpustakaan/model/domain"
	"perpustakaan/model/web/request"
)

type UserService interface {
	FindAllUser() ([]domain.User, error)
	CreateUser(input request.UserRequest) (domain.User, error)
	FindOneUser(Id int) (domain.User, error)
	UpdateUser(Id int, input request.UserUpdateRequest) (domain.User, error)
	DeleteUser(Id int) (domain.User, error)
}
