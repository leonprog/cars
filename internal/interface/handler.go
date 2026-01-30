package _interface

import "net/http"

type Handler interface {
	GetCars(w http.ResponseWriter, r *http.Request)
	FindCarsById(w http.ResponseWriter, r *http.Request)
	AddCar(w http.ResponseWriter, r *http.Request)
	DeleteCar(w http.ResponseWriter, r *http.Request)
}
