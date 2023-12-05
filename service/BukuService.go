package service

import (
	"perpustakaan/model/domain"
	"perpustakaan/model/web/request"
)

type BukuService interface {
	BukuFindAll() ([]domain.Buku, error)
	BukuFindOne(Id int) (domain.Buku, error)
	BukuCreate(input request.BukuRequest) (domain.Buku, error)
	BukuUpdate(Id int, input request.BukuRequest) (domain.Buku, error)
	BukuDelete(Id int) (domain.Buku, error)
}
