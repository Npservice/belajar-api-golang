package ServiceImpl

import (
	"perpustakaan/model/domain"
	"perpustakaan/model/web/request"
	"perpustakaan/repository"
)

type bukuRespository struct {
	bukuRespository repository.BukuRepository
}

func NewBukuService(bukuRepository repository.BukuRepository) *bukuRespository {
	return &bukuRespository{bukuRepository}
}

func (r *bukuRespository) BukuFindAll() ([]domain.Buku, error) {
	var buku []domain.Buku
	buku, err := r.bukuRespository.FindAll()
	if err != nil {
		return buku, err
	}
	return buku, nil
}
func (r *bukuRespository) BukuFindOne(Id int) (domain.Buku, error) {
	buku, err := r.bukuRespository.FindOne(Id)
	if err != nil {
		return buku, err
	}
	return buku, nil
}
func (r *bukuRespository) BukuCreate(input request.BukuRequest) (domain.Buku, error) {
	var buku domain.Buku
	buku.Judul = input.Judul
	buku.Penerbit = input.Penerbit
	buku.Pengarang = input.Pengarang
	buku.JumlahStok = input.JumlahStok
	buku.TahunTerbit = input.TahunTerbit

	Buku, err := r.bukuRespository.Create(buku)
	if err != nil {
		return Buku, err
	}
	return Buku, nil

}
func (r *bukuRespository) BukuUpdate(Id int, input request.BukuRequest) (domain.Buku, error) {

	var buku domain.Buku
	buku.Judul = input.Judul
	buku.Penerbit = input.Penerbit
	buku.Pengarang = input.Pengarang
	buku.JumlahStok = input.JumlahStok
	buku.TahunTerbit = input.TahunTerbit

	Buku, err := r.bukuRespository.Update(Id, buku)
	if err != nil {
		return Buku, err
	}
	return Buku, nil
}

func (r *bukuRespository) BukuDelete(Id int) (domain.Buku, error) {
	var buku domain.Buku
	Buku, err := r.bukuRespository.Delete(Id, buku)
	if err != nil {
		return Buku, err
	}
	return Buku, err
}
