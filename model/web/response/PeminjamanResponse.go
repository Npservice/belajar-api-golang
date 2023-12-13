package response

import (
	"perpustakaan/model/domain"
	"time"
)

type peminjamanResponse struct {
	Id                  int    `json:"id"`
	AnggotaId           int    `json:"anggota_id"`
	BukuId              int    `json:"buku_id"`
	NamaAnggota         string `json:"nama_anggota"`
	Alamat              string `json:"alamat"`
	Telephone           string `json:"telphone"`
	JudulBuku           string `json:"judul_buku"`
	Pengarang           string `json:"pengarang"`
	Penerbit            string `json:"penerbit"`
	TanggalPeminjaman   string `json:"tanggal_peminjaman"`
	TanggalPengembalian string `json:"tanggal_pengembalian"`
}

func tanggalFormater(times time.Time) string {
	dateString := times.Format("2006-01-02")
	return dateString
}

func ResponsePeminjaman(peminjaman domain.Peminjaman) peminjamanResponse {
	var response peminjamanResponse
	response.Id = peminjaman.ID
	response.AnggotaId = peminjaman.AnggotaID
	response.BukuId = peminjaman.BukuID
	response.NamaAnggota = peminjaman.Anggota.NamaDepan
	response.Alamat = peminjaman.Anggota.Alamat
	response.Telephone = peminjaman.Anggota.Telephone
	response.JudulBuku = peminjaman.Buku.Judul
	response.Pengarang = peminjaman.Buku.Pengarang
	response.Penerbit = peminjaman.Buku.Penerbit
	response.TanggalPeminjaman = tanggalFormater(peminjaman.TanggalPeminjaman)
	response.TanggalPengembalian = tanggalFormater(peminjaman.TanggalPengembalian)
	return response
}

func ResponsePeminjamans(peminjamans []domain.Peminjaman) []peminjamanResponse {
	var response []peminjamanResponse
	for _, peminjaman := range peminjamans {
		formatter := ResponsePeminjaman(peminjaman)
		response = append(response, formatter)
	}
	return response
}
