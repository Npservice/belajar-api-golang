package RepositoryImpl

import (
	"gorm.io/gorm"
	"perpustakaan/model/domain"
)

type database struct {
	db *gorm.DB
}

func NewAnggotaRepository(db *gorm.DB) *database {
	return &database{db}
}

func (d *database) FindAll() ([]domain.Anggota, error) {
	var anggota []domain.Anggota
	err := d.db.Find(&anggota).Error
	if err != nil {
		return nil, err
	}
	return anggota, nil
}