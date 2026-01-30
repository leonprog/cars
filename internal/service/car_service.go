package service

import (
	"api/internal/dto"
	_interface "api/internal/interface"
	"api/internal/models"
)

type CarService struct {
	rep _interface.Repository
}

func New(rep _interface.Repository) _interface.Service {
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

func (s *CarService) FindCarById(id int) (models.Car, error) {
	car, err := s.rep.FindById(id)

	return car, err
}

func (s *CarService) CreateCar(createCarDto dto.CarCreateDto) (int, error) {
	id, err := s.rep.Create(createCarDto)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *CarService) Delete(id int) error {
	return s.rep.Delete(id)
}
