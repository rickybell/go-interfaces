package config

import (
	"github.com/rickybell/go-interfaces/app/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:@localhost:5432/go-interfaces"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate((&entities.User{}))
	DB = db
}
