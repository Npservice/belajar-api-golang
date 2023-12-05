package repository

import "perpustakaan/model/domain"

type UserRepository interface {
	FindAll() ([]domain.User, error)
	Create(user domain.User) (domain.User, error)
	FindOne(Id int, user domain.User) (domain.User, error)
	Update(Id int, user domain.User) (domain.User, error)
	Delete(Id int, user domain.User) (domain.User, error)
}
