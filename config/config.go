package config

import (
	"os"

	"github.com/fakhrizalmus/tabungango/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DB")
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{})
	DB = db
}
