package RepositoryImpl

import (
	"gorm.io/gorm"
	"perpustakaan/model/domain"
)

type peminjamRepostory struct {
	db *gorm.DB
}

func NewPeminjamanRepository(db *gorm.DB) *peminjamRepostory {
	return &peminjamRepostory{db}
}

func (r *peminjamRepostory) FindAll() ([]domain.Peminjaman, error) {
	var peminjaman []domain.Peminjaman
	err := r.db.Preload("Buku").Preload("Anggota").Find(&peminjaman).Error
	if err != nil {
		return peminjaman, err
	}
	return peminjaman, nil
}
