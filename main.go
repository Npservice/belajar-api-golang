package main

import (
	"perpustakaan/app"
	ControllerImpl "perpustakaan/controller/impl"
	RepositoryImpl "perpustakaan/repository/impl"
	ServiceImpl "perpustakaan/service/impl"
)

func main() {
	db := app.NewDB()

	UserRepository := RepositoryImpl.NewUserRepository(db)
	UserService := ServiceImpl.NewUserService(UserRepository)
	User := ControllerImpl.NewUserController(UserService)

	BukuRepository := RepositoryImpl.NewBukuRepository(db)
	BukuService := ServiceImpl.NewBukuService(BukuRepository)
	Buku := ControllerImpl.NewBukuController(BukuService)

	AnggotaRepository := RepositoryImpl.NewAnggotaRepository(db)
	AnggotaService := ServiceImpl.NewAnggotaService(AnggotaRepository)
	Anggota := ControllerImpl.NewAnggotaController(AnggotaService)

	r := &app.Route{
		User,
		Buku,
		Anggota,
	}
	route := app.NewRoute(r)

	route.Run(":9000")

}
