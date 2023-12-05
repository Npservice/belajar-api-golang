package domain

type Terlambat struct {
	IdKeterlambatan int
	IdPeminjaman    Peminjaman
	Denda           int
}
