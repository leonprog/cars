package handler

import (
	"api/internal/dto"
	"api/internal/handler/request"
	_interface "api/internal/interface"
	"api/pkg/res"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

type CarHandler struct {
	service _interface.Service
}

func New(router *http.ServeMux, service _interface.Service) {
	handler := &CarHandler{
		service: service,
	}

	router.HandleFunc("POST /cars/add", handler.AddCar)
	router.HandleFunc("GET /cars", handler.GetCars)
	router.HandleFunc("GET /cars/{id}", handler.FindCarsById)
	router.HandleFunc("DELETE /cars/{id}", handler.DeleteCar)
}

func (h *CarHandler) GetCars(w http.ResponseWriter, r *http.Request) {
	cars, err := h.service.GetCars()
	if err != nil {
		res.JsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	res.JsonRes(w, cars, http.StatusOK)
}

func (h *CarHandler) FindCarsById(w http.ResponseWriter, r *http.Request) {
	idParse := r.PathValue("id")

	id, err := strconv.Atoi(idParse)
	if err != nil {
		res.JsonError(w, err.Error(), http.StatusBadRequest)

		return
	}

	car, err := h.service.FindCarById(id)
	if err != nil {
		res.JsonError(w, err.Error(), http.StatusBadRequest)

		return
	}

	res.JsonRes(w, car, http.StatusOK)
}

func (h *CarHandler) AddCar(w http.ResponseWriter, r *http.Request) {
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

	carCreateDto := dto.CarCreateDto{
		Mark:       carPayload.Mark,
		Model:      carPayload.Model,
		OwnerCount: carPayload.OwnerCount,
		Price:      carPayload.Price,
		Currency:   carPayload.Currency,
		Options:    carPayload.Options,
	}

	result, err := h.service.CreateCar(carCreateDto)
	if err != nil {
		res.JsonError(w, err.Error(), http.StatusBadRequest)

		return
	}

	res.JsonRes(w, result, 200)
}

func (h *CarHandler) DeleteCar(w http.ResponseWriter, r *http.Request) {
	idParse := r.PathValue("id")

	id, err := strconv.Atoi(idParse)
	if err != nil {
		res.JsonError(w, err.Error(), http.StatusBadRequest)

		return
	}

	err = h.service.Delete(id)
	if err != nil {
		res.JsonError(w, err.Error(), http.StatusBadRequest)

		return
	}

	res.JsonRes(w, nil, http.StatusOK)
}
