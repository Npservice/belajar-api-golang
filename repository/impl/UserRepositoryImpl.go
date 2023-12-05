package RepositoryImpl

import (
	"perpustakaan/model/domain"

	"gorm.io/gorm"
)

type userRespository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRespository {
	return &userRespository{db}
}

func (r *userRespository) FindAll() ([]domain.User, error) {
	var users []domain.User

	err := r.db.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRespository) Create(user domain.User) (domain.User, error) {

	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRespository) FindOne(Id int, user domain.User) (domain.User, error) {

	err := r.db.Where("id=?", Id).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRespository) Update(Id int, user domain.User) (domain.User, error) {

	err := r.db.Where("id=?", Id).Updates(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil

}

func (r *userRespository) Delete(Id int, user domain.User) (domain.User, error) {
	err := r.db.Where("id=?", Id).Delete(&user).Error
	if err != nil {
		return user, err
	}
	return user, err
}
