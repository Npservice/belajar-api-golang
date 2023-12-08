package repository

import "perpustakaan/model/domain"

type AnggotaRepository interface {
	FindAll() ([]domain.Anggota, error)
	FindOne(Id int) (domain.Anggota, error)
	Create(anggota domain.Anggota) (domain.Anggota, error)
	//Update(Id int, anggota domain.Anggota) (domain.Anggota, error)
	//Delete(Id int) (domain.Anggota, error)
}
