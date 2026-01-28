package repository

import (
	"api/internal/handler/request"
	"api/internal/models"
	"api/pkg/db"
)

type CarRepository struct {
	Db *db.Db
}

func NewRepository(db *db.Db) *CarRepository {
	return &CarRepository{
		Db: db,
	}
}

func (r *CarRepository) Get() ([]models.Car, error) {
	var cars []models.Car

	res := r.Db.
		Db.
		Table("cars").
		Scan(&cars)

	if res.Error != nil {
		return nil, res.Error
	}

	return cars, nil
}

func (r *CarRepository) FindById(id int) (*models.Car, error) {
	var car models.Car

	res := r.Db.
		Db.
		Table("cars").
		First(&car)

	if res.Error != nil {
		return nil, res.Error
	}

	return &car, nil

}

func (r *CarRepository) Create(carRequest request.CreateCarRequest) (*models.Car, error) {
	car := &models.Car{
		Mark:       carRequest.Mark,
		ModelCar:   carRequest.Model,
		OwnerCount: carRequest.OwnerCount,
		Price:      carRequest.Price,
		Currency:   carRequest.Currency,
		Options:    carRequest.Options,
	}

	res := r.Db.Db.Create(car)

	if res.Error != nil {
		return nil, res.Error
	}

	return car, nil
}

func (r *CarRepository) Delete(id int) error {
	res := r.Db.Db.Table("cars").Unscoped().Delete(&models.Car{}, id)

	if res.Error != nil {
		return res.Error
	}

	return nil
}
