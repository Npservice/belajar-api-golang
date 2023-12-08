package app

import (
	"github.com/gin-gonic/gin"
	"perpustakaan/controller"
	"perpustakaan/exception"
)

type Route struct {
	UserController    controller.UserController
	BukuController    controller.BukuController
	AnggotaController controller.AnggotaController
}

func NewRoute(r *Route) *gin.Engine {
	route := gin.Default()
	gin.SetMode(gin.DebugMode)
	route.Use(exception.NewErrorHandler)

	route.GET("/user", r.UserController.GetAllUser)
	route.POST("/user/create", r.UserController.CreateUser)
	route.GET("/user/find/:id", r.UserController.FindOneUser)
	route.PATCH("/user/update/:id", r.UserController.UpdateUser)
	route.DELETE("/user/destroy/:id", r.UserController.DeleteUser)

	route.GET("/buku", r.BukuController.FindAllBuku)
	route.GET("/buku/:id", r.BukuController.FindOneBuku)
	route.POST("/buku/create", r.BukuController.CreateBuku)
	route.PATCH("/buku/update/:id", r.BukuController.UpdateBuku)
	route.DELETE("/buku/destroy/:id", r.BukuController.DeleteBuku)

	route.GET("/anggota", r.AnggotaController.FindAllAnggota)
	route.GET("/anggota/:id", r.AnggotaController.FindOneAnggota)

	return route
}
