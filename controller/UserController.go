package controller

import "github.com/gin-gonic/gin"

type UserController interface {
	GetAllUser(c *gin.Context)
	CreateUser(c *gin.Context)
	FindOneUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}
