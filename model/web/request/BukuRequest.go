package request

type BukuRequest struct {
	Judul       string `json:"judul" binding:"required"`
	Pengarang   string `json:"pengarang"  binding:"required"`
	Penerbit    string `json:"penerbit"  binding:"required"`
	TahunTerbit int    `json:"tahun_terbit"  binding:"required"`
	JumlahStok  int    `json:"jumlah_stok"  binding:"required"`
}
