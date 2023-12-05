package controller

import "github.com/gin-gonic/gin"

type AnggotaController interface {
	FindAllAnggota(c *gin.Context)
}
