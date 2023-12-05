package domain

import (
	"time"
)

type Peminjaman struct {
	IDPeminjaman        int
	IdAnggota           Anggota
	IdBuku              Buku
	TanggalPeminjaman   time.Time
	TanggalPengembalian time.Time
}
