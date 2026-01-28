package handler

import (
	"api/internal/handler/request"
	"api/internal/service"
	"api/pkg/res"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

type CarHandler struct {
	service *service.CarService
}

func NewHandler(router *http.ServeMux, service *service.CarService) {
	handler := &CarHandler{
		service: service,
	}

	router.HandleFunc("POST /cars/add", handler.AddCar())
	router.HandleFunc("GET /cars", handler.GetCars())
	router.HandleFunc("GET /cars/{id}", handler.FindCarsById())
	router.HandleFunc("DELETE /cars/{id}", handler.DeleteCar())
}

func (h *CarHandler) GetCars() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cars, err := h.service.GetCars()

		if err != nil {
			res.JsonError(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.JsonRes(w, cars, 200)
	}
}

func (h *CarHandler) FindCarsById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParse := r.PathValue("id")

		id, err := strconv.Atoi(idParse)

		car, err := h.service.FindCarById(id)

		if err != nil {
			res.JsonError(w, err.Error(), http.StatusBadRequest)

			return
		}

		res.JsonRes(w, car, 200)
	}
}

func (h *CarHandler) AddCar() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var carPayload request.CreateCarRequest

		err := json.NewDecoder(r.Body).Decode(&carPayload)

		if err != nil {
			res.JsonError(w, err.Error(), http.StatusBadRequest)

			return
		}

		validate := validator.New()

		err = validate.Struct(&carPayload)

		if err != nil {
			res.JsonError(w, err.Error(), http.StatusBadRequest)

			return
		}

		result, err := h.service.CreateCar(carPayload)

		if err != nil {
			res.JsonError(w, err.Error(), http.StatusBadRequest)

			return
		}

		res.JsonRes(w, result, 200)
	}
}

func (h *CarHandler) DeleteCar() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParse := r.PathValue("id")

		id, err := strconv.Atoi(idParse)

		err = h.service.Delete(id)

		if err != nil {
			res.JsonError(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.JsonRes(w, nil, 200)
	}
}
