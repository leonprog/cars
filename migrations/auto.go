package main

import (
	"api/internal/config"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"log"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	c := config.New()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s  sslmode=disable",
		c.DbConfig.Host, c.DbConfig.Port, c.DbConfig.Username, c.DbConfig.Password, c.DbConfig.Db,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	goose.SetDialect("postgres")

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}
}
