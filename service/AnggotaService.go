package service

import "perpustakaan/model/domain"

type AnggotaService interface {
	FindAllAnggota() ([]domain.Anggota, error)
	FindOneAnggota(Id int) (domain.Anggota, error)
}
