package db

import (
	"api/internal/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	Db gorm.DB
}

func NewDb(c *config.Config) *Db {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", c.DbConfig.Host, c.DbConfig.Username, c.DbConfig.Password, c.DbConfig.Db, c.DbConfig.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return &Db{
		Db: *db,
	}
}
