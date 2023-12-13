package service

import "perpustakaan/model/domain"

type PeminjamanService interface {
	FindAllPeminjaman() ([]domain.Peminjaman, error)
}
