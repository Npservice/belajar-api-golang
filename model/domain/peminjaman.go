package domain

import (
	"time"
)

type Peminjaman struct {
	ID                  int
	AnggotaID           int
	Anggota             Anggota
	BukuID              int
	Buku                Buku
	TanggalPeminjaman   time.Time
	TanggalPengembalian time.Time
}
