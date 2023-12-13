package ControllerImpl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"perpustakaan/helpers"
	"perpustakaan/model/web/response"
	"perpustakaan/service"
)

type peminjamanService struct {
	service service.PeminjamanService
}

func NewPeminjamanController(service service.PeminjamanService) *peminjamanService {
	return &peminjamanService{service}
}

func (r *peminjamanService) PeminjamanFindALl(c *gin.Context) {
	peminjaman, err := r.service.FindAllPeminjaman()

	if err != nil {
		c.Error(err)
		return
	}
	respond := helpers.ApiResponse("Peminjaman Found", http.StatusOK, "success", response.ResponsePeminjamans(peminjaman))
	c.JSON(http.StatusOK, respond)
	return
}
