package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type TourHandler struct {
	TourService *service.TourService
}

func NewTourHandler(db *gorm.DB) *TourHandler {
	tourService := service.NewTourService(db)
	return &TourHandler{TourService: tourService}
}

func (h *TourHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/tours/{id}", h.Get).Methods("GET")
	router.HandleFunc("/tours", h.Create).Methods("POST")
}

func (handler *TourHandler) Get(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	idString := params["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println("Error parsing ID:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("Tour with ID: %d", id)
	tour, err := handler.TourService.Find(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(tour)
}

func (handler *TourHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var tour model.Tour
	err := json.NewDecoder(req.Body).Decode(&tour)
	if err != nil {
		log.Println("Error while parsing JSON:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	tour.ID = int(tour.ID)

	err = handler.TourService.Create(&tour)
	if err != nil {
		log.Println("Error while creating new tour:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
