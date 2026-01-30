package _interface

import (
	"api/internal/dto"
	"api/internal/models"
)

type Service interface {
	GetCars() ([]models.Car, error)
	FindCarById(id int) (models.Car, error)
	CreateCar(createCarDto dto.CarCreateDto) (int, error)
	Delete(id int) error
}
