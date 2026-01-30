package _interface

import (
	"api/internal/dto"
	"api/internal/models"
)

type Repository interface {
	Get() ([]models.Car, error)
	FindById(id int) (models.Car, error)
	Create(dto dto.CarCreateDto) (int, error)
	Delete(id int) error
}
