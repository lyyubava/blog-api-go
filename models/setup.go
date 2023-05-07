package models

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DATABASE_CONNECTION_STRING")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Fail to connect to db")
	}
	err = database.AutoMigrate(&User{}, &Post{}, &Comment{})

	if err != nil {
		return
	}
	DB = database
}
