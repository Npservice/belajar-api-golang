package ServiceImpl

import (
	"perpustakaan/model/domain"
	"perpustakaan/repository"
)

type anggotaRepository struct {
	repository repository.AnggotaRepository
}

func NewAnggotaService(repository repository.AnggotaRepository) *anggotaRepository {
	return &anggotaRepository{repository}
}

func (r *anggotaRepository) FindAllAnggota() ([]domain.Anggota, error) {
	anggota, err := r.repository.FindAll()
	if err != nil {
		return anggota, err
	}
	return anggota, nil
}
func (r *anggotaRepository) FindOneAnggota(Id int) (domain.Anggota, error) {
	anggota, err := r.repository.FindOne(Id)
	if err != nil {
		return anggota, err
	}
	return anggota, nil
}
