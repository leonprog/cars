package main

import (
	config2 "api/internal/config"
	"api/internal/models"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err.Error())
	}

	c := config2.NewConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", c.DbConfig.Host, c.DbConfig.Username, c.DbConfig.Password, c.DbConfig.Db, c.DbConfig.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Car{})
}
