package response

import "perpustakaan/model/domain"

type bukuResponse struct {
	Id          int    `json:"id"`
	Judul       string `json:"judul"`
	Pengarang   string `json:"pengarang"`
	Penerbit    string `json:"penerbit"`
	TahunTerbit int    `json:"tahun_terbit"`
	JumlahStok  int    `json:"jumlah_stok"`
}

func ResponseBuku(Buku domain.Buku) bukuResponse {
	formatter := bukuResponse{}
	formatter.Id = Buku.Id
	formatter.Judul = Buku.Judul
	formatter.Pengarang = Buku.Pengarang
	formatter.Penerbit = Buku.Penerbit
	formatter.TahunTerbit = Buku.TahunTerbit
	formatter.JumlahStok = Buku.JumlahStok
	return formatter
}

func ResponsBukus(Buku []domain.Buku) []bukuResponse {

	var bukuResponses []bukuResponse
	for _, bukus := range Buku {
		formatter := ResponseBuku(bukus)
		bukuResponses = append(bukuResponses, formatter)
	}
	return bukuResponses

}
