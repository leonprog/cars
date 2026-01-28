package app

import (
	"api/internal/config"
	"api/internal/handler"
	"api/internal/repository"
	"api/internal/service"
	"api/pkg/db"
	"net/http"
)

func InitApp() {

	router := http.NewServeMux()

	newConfig := config.NewConfig()
	newDb := db.NewDb(newConfig)

	// Repositories
	carRepository := repository.NewRepository(newDb)

	// Services
	carService := service.NewService(carRepository)

	// Handlers
	handler.NewHandler(router, carService)

	http.ListenAndServe(
		":8090",
		router,
	)
}
