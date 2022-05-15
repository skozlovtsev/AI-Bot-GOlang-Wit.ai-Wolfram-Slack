package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

var (
	db *gorm.DB
)

func Connect(){
	//Подключение к базе данных sqlite
	database, err := gorm.Open(sqlite.Open("users_test"), &gorm.Config{})
	if err != nil{
		panic(err)
	}
	db = database
}

func GetDB() *gorm.DB {
	return db
}