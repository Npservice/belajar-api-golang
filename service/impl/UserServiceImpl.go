package ServiceImpl

import (
	"perpustakaan/model/domain"
	"perpustakaan/model/web/request"
	"perpustakaan/repository"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) FindAllUser() ([]domain.User, error) {

	users, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil

}

func (s *userService) CreateUser(input request.UserRequest) (domain.User, error) {

	user := domain.User{}
	user.Name = input.Name
	user.Username = input.Username
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)

	newUser, err := s.repository.Create(user)
	if err != nil {
		return user, err
	}
	return newUser, nil

}
func (s *userService) FindOneUser(Id int) (domain.User, error) {
	DomainUser := domain.User{}
	User, err := s.repository.FindOne(Id, DomainUser)
	if err != nil {
		return User, nil
	}
	return User, nil
}

func (s *userService) UpdateUser(Id int, input request.UserUpdateRequest) (domain.User, error) {
	User := domain.User{}
	User.Name = input.Name
	User.Username = input.Username
	if input.Password != "" {
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
		if err != nil {
			return User, err
		}
		User.Password = string(passwordHash)
	}
	updateUser, err := s.repository.Update(Id, User)

	if err != nil {
		return User, err
	}

	return updateUser, nil
}

func (s *userService) DeleteUser(Id int) (domain.User, error) {
	User := domain.User{}
	user, err := s.repository.Delete(Id, User)
	if err != nil {
		return User, err
	}
	return user, nil

}
