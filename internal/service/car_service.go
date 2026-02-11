package service

import (
	"api/internal/dto"
	"api/internal/models"
	"context"
)

type Repository interface {
	Get(ctx context.Context) ([]models.Car, error)
	FindById(ctx context.Context, id int) (models.Car, error)
	Create(ctx context.Context, dto dto.CarCreateDto) (int, error)
	Delete(ctx context.Context, id int) error
}

type CarService struct {
	rep Repository
}

func New(rep Repository) *CarService {
	return &CarService{
		rep: rep,
	}
}

func (s *CarService) GetCars(ctx context.Context) ([]models.Car, error) {
	cars, err := s.rep.Get(ctx)
	if err != nil {
		return nil, err
	}

	return cars, nil
}

func (s *CarService) FindCarById(ctx context.Context, id int) (models.Car, error) {
	car, err := s.rep.FindById(ctx, id)

	return car, err
}

func (s *CarService) CreateCar(ctx context.Context, createCarDto dto.CarCreateDto) (int, error) {
	id, err := s.rep.Create(ctx, createCarDto)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *CarService) Delete(ctx context.Context, id int) error {
	return s.rep.Delete(ctx, id)
}
