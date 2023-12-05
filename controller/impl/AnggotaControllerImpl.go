package ControllerImpl

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"perpustakaan/helpers"
	"perpustakaan/model/web/response"
	"perpustakaan/service"
)

type anggotaService struct {
	service service.AnggotaService
}

func NewAnggotaController(service service.AnggotaService) *anggotaService {
	return &anggotaService{service}
}

func (s *anggotaService) FindAllAnggota(c *gin.Context) {
	anggota, err := s.service.FindAllAnggota()
	if err != nil {
		c.Error(err)
		return
	}
	if len(anggota) == 0 {
		c.Error(errors.New("not Found"))
		return
	}
	formatter := response.ResponseAnggotas(anggota)
	respond := helpers.ApiResponse("Anggota Found", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, respond)
}
