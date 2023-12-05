package repository

import "perpustakaan/model/domain"

type BukuRepository interface {
	FindAll() ([]domain.Buku, error)
	FindOne(Id int) (domain.Buku, error)
	Create(buku domain.Buku) (domain.Buku, error)
	Update(Id int, buku domain.Buku) (domain.Buku, error)
	Delete(Id int, buku domain.Buku) (domain.Buku, error)
}
