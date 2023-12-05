package ControllerImpl

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"perpustakaan/helpers"
	"perpustakaan/model/web/request"
	"perpustakaan/model/web/response"
	"perpustakaan/service"
	"strconv"
)

type bukuSevice struct {
	bukuSevice service.BukuService
}

func NewBukuController(service service.BukuService) *bukuSevice {
	return &bukuSevice{service}
}

func (s *bukuSevice) FindAllBuku(c *gin.Context) {
	buku, err := s.bukuSevice.BukuFindAll()
	if err != nil {
		c.Error(err)
		return
	}
	if len(buku) == 0 {
		c.Error(errors.New("not Found"))
		return
	}
	respond := helpers.ApiResponse("Buku Found", http.StatusOK, "success", response.ResponsBukus(buku))

	c.JSON(http.StatusOK, respond)
	return
}
func (s *bukuSevice) FindOneBuku(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.Error(err)
		return
	}
	buku, err := s.bukuSevice.BukuFindOne(id)
	if err != nil {
		c.Error(err)
		return
	}
	if buku.Id == 0 {
		c.Error(errors.New("not Found"))
		return
	}
	respond := helpers.ApiResponse("Buku Found", http.StatusOK, "success", response.ResponseBuku(buku))

	c.JSON(http.StatusOK, respond)
	return

}
func (s *bukuSevice) CreateBuku(c *gin.Context) {
	var input request.BukuRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		validator := helpers.ValidatorHandlerError(err)
		respond := helpers.ApiResponse("Buku Not Valid", http.StatusUnprocessableEntity, "faild", gin.H{"error": validator})
		c.JSON(http.StatusUnprocessableEntity, respond)
		return
	}
	Buku, err := s.bukuSevice.BukuCreate(input)
	if err != nil {
		c.Error(err)
		return
	}
	respond := helpers.ApiResponse("Buku Create Successfully", http.StatusOK, "success", response.ResponseBuku(Buku))
	c.JSON(http.StatusOK, respond)
	return
}
func (s *bukuSevice) UpdateBuku(c *gin.Context) {
	param := c.Param("id")
	Id, err := strconv.Atoi(param)
	if err != nil {
		c.Error(err)
		return
	}
	var input request.BukuRequest
	validation := c.ShouldBindJSON(&input)
	if validation != nil {
		validator := helpers.ValidatorHandlerError(validation)
		respond := helpers.ApiResponse("Buku Not Valid", http.StatusUnprocessableEntity, "faild", gin.H{"error": validator})
		c.JSON(http.StatusUnprocessableEntity, respond)
		return
	}

	Buku, err := s.bukuSevice.BukuUpdate(Id, input)
	if err != nil {
		c.Error(err)
		return
	}
	respond := helpers.ApiResponse("Buku Update Successfully", http.StatusOK, "success", response.ResponseBuku(Buku))
	c.JSON(http.StatusOK, respond)
	return
}
func (s *bukuSevice) DeleteBuku(c *gin.Context) {
	param := c.Param("id")
	Id, err := strconv.Atoi(param)
	if err != nil {
		c.Error(err)
		return
	}
	Buku, err := s.bukuSevice.BukuDelete(Id)
	if err != nil {
		c.Error(err)
		return
	}
	respond := helpers.ApiResponse("Buku Delete Successfully", http.StatusOK, "success", response.ResponseBuku(Buku))
	c.JSON(http.StatusOK, respond)
	return

}
