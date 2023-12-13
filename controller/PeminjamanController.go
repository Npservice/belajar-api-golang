package controller

import "github.com/gin-gonic/gin"

type PeminjamanController interface {
	PeminjamanFindALl(c *gin.Context)
}
