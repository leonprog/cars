package repository

import (
	"api/internal/dto"
	"api/internal/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type CarRepository struct {
	db *pgx.Conn
}

func New(db *pgx.Conn) *CarRepository {
	return &CarRepository{
		db: db,
	}
}

func (r *CarRepository) Get(ctx context.Context) ([]models.Car, error) {
	var cars []models.Car

	res, err := r.db.Query(ctx, "SELECT id, mark, model_car, owner_count, price, currency, options FROM cars;")
	if err != nil {
		return nil, err
	}
	defer res.Close()

	for res.Next() {
		var car models.Car

		err = res.Scan(
			&car.Id,
			&car.Mark,
			&car.ModelCar,
			&car.OwnerCount,
			&car.Price,
			&car.Currency,
			&car.Options,
		)
		if err != nil {
			fmt.Println(err)
			continue
		}

		cars = append(cars, car)
	}

	return cars, nil
}

func (r *CarRepository) FindById(ctx context.Context, id int) (models.Car, error) {
	var car models.Car

	err := r.db.
		QueryRow(ctx,
			"SELECT id, mark, model_car, owner_count, price, currency, options FROM cars WHERE id=$1 LIMIT 1;",
			id,
		).Scan(
		&car.Id,
		&car.Mark,
		&car.ModelCar,
		&car.OwnerCount,
		&car.Price,
		&car.Currency,
		&car.Options,
	)

	return car, err
}

func (r *CarRepository) Create(ctx context.Context, dto dto.CarCreateDto) (int, error) {
	var id int

	query := "INSERT INTO cars (mark, model_car, owner_count,price,currency, options) VALUES($1,$2,$3,$4,$5,$6) RETURNING Id;"

	err := r.db.QueryRow(ctx, query, dto.Mark, dto.Model, dto.OwnerCount, dto.Price, dto.Currency, dto.Options).Scan(&id)

	return id, err
}

func (r *CarRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, "DELETE FROM cars WHERE id=$1;", id)

	return err
}
