package RepositoryImpl

import (
	"gorm.io/gorm"
	"perpustakaan/model/domain"
)

type bukuRepository struct {
	db *gorm.DB
}

func NewBukuRepository(db *gorm.DB) *bukuRepository {
	return &bukuRepository{db}
}

func (r *bukuRepository) FindAll() ([]domain.Buku, error) {
	var buku []domain.Buku

	err := r.db.Find(&buku).Error
	if err != nil {
		return nil, err
	}
	return buku, nil

}

func (r *bukuRepository) FindOne(Id int) (domain.Buku, error) {

	var buku domain.Buku
	err := r.db.Where("id=?", Id).Find(&buku).Error

	if err != nil {
		return buku, err
	}
	return buku, nil
}
func (r *bukuRepository) Create(buku domain.Buku) (domain.Buku, error) {
	err := r.db.Create(&buku).Error
	if err != nil {
		return buku, err
	}
	return buku, nil
}

func (r *bukuRepository) Update(Id int, buku domain.Buku) (domain.Buku, error) {
	err := r.db.Where("id=?", Id).Updates(&buku).Error
	if err != nil {
		return buku, nil
	}
	return buku, nil
}
func (r *bukuRepository) Delete(Id int, buku domain.Buku) (domain.Buku, error) {
	err := r.db.Where("id=?", Id).Delete(&buku).Error
	if err != nil {
		return buku, err
	}
	return buku, nil

}
