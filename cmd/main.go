package main

import (
	"api/internal/config"
	"api/internal/handler"
	"api/internal/repository"
	"api/internal/service"
	"api/pkg/db"
	"context"
	"net/http"
)

func main() {
	ctx := context.Background()
	router := http.NewServeMux()

	newConfig := config.New()
	newDb := db.New(ctx, newConfig)
	defer newDb.Close(ctx)

	// Repositories
	carRepository := repository.New(ctx, newDb)

	// Services
	carService := service.New(carRepository)

	// Handlers
	handler.New(router, carService)

	http.ListenAndServe(
		":8090",
		router,
	)
}
