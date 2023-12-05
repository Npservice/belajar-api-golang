package app

import (
	"perpustakaan/helpers"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/perpustakaan?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helpers.PanicIfError(err)
	return db
}
