package service

import (
	"api/internal/handler/request"
	"api/internal/models"
	"api/internal/repository"
)

type CarService struct {
	rep *repository.CarRepository
}

func NewService(rep *repository.CarRepository) *CarService {
	return &CarService{
		rep: rep,
	}
}

func (s *CarService) GetCars() ([]models.Car, error) {
	cars, err := s.rep.Get()

	if err != nil {
		return nil, err
	}

	return cars, nil
}

func (s *CarService) FindCarById(id int) (*models.Car, error) {
	car, err := s.rep.FindById(id)

	if err != nil {
		return nil, err
	}

	return car, nil
}

func (s *CarService) CreateCar(createCarReq request.CreateCarRequest) (int, error) {
	car, err := s.rep.Create(createCarReq)

	if err != nil {
		return 0, err
	}

	return int(car.ID), nil
}

func (s *CarService) Delete(id int) error {
	err := s.rep.Delete(id)

	return err
}
