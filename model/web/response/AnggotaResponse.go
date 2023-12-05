package response

import "perpustakaan/model/domain"

type AnggotaResponse struct {
	Id           int    `json:"id"`
	NamaDepan    string `json:"nama_depan"`
	NamaBelakang string `json:"nama_belakang"`
	Alamat       string `json:"alamat"`
	Telephone    string `json:"telephone"`
}

func ResponseAnggota(anggota domain.Anggota) AnggotaResponse {

	var Response AnggotaResponse

	Response.Id = anggota.Id
	Response.NamaDepan = anggota.NamaDepan
	Response.NamaBelakang = anggota.NamaBelakang
	Response.Alamat = anggota.Alamat
	Response.Telephone = anggota.Telephone
	return Response
}

func ResponseAnggotas(anggotas []domain.Anggota) []AnggotaResponse {
	var Responses []AnggotaResponse

	for _, anggota := range anggotas {
		var formatter = ResponseAnggota(anggota)
		Responses = append(Responses, formatter)
	}
	return Responses

}
