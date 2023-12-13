package ServiceImpl

import (
	"perpustakaan/model/domain"
	"perpustakaan/repository"
)

type peminjamanRepository struct {
	repository repository.PeminjamanRepository
}

func NewPeminjamanService(repository repository.PeminjamanRepository) *peminjamanRepository {
	return &peminjamanRepository{repository}
}

func (r *peminjamanRepository) FindAllPeminjaman() ([]domain.Peminjaman, error) {
	peminjam, err := r.repository.FindAll()
	if err != nil {
		return peminjam, err
	}
	return peminjam, nil
}
