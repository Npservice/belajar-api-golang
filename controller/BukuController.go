package controller

import "github.com/gin-gonic/gin"

type BukuController interface {
	FindAllBuku(c *gin.Context)
	FindOneBuku(c *gin.Context)
	CreateBuku(c *gin.Context)
	UpdateBuku(c *gin.Context)
	DeleteBuku(c *gin.Context)
}
