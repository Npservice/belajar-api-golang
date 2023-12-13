package repository

import "perpustakaan/model/domain"

type PeminjamanRepository interface {
	FindAll() ([]domain.Peminjaman, error)
}
