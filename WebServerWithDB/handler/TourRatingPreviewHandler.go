package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

type TourRatingPreviewHandler struct {
	TourRatingPreviewService *service.TourRatingPreviewService
}

func NewTourRatingPreviewHandler(db *gorm.DB) *TourRatingPreviewHandler {
	tourRatingPreviewService := service.NewTourRatingPreviewService(db)
	return &TourRatingPreviewHandler{TourRatingPreviewService: tourRatingPreviewService}
}

func (h *TourRatingPreviewHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/tourRatingPreviews/{id}", h.Get).Methods("GET")
	router.HandleFunc("/tourRatingPreviews", h.Create).Methods("POST")
}

func (handler *TourRatingPreviewHandler) Get(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	idString := params["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println("Error parsing ID:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("Checkpoint with ID: %d", id)
	tourRatingPreview, err := handler.TourRatingPreviewService.Find(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(tourRatingPreview)
}

func (handler *TourRatingPreviewHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var tourRatingPreview model.TourRatingPreview
	err := json.NewDecoder(req.Body).Decode(&tourRatingPreview)
	if err != nil {
		log.Println("Error while parsing JSON:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if tourRatingPreview.TourID == 0 || tourRatingPreview.TouristID == 0 || tourRatingPreview.Rating == 0 || tourRatingPreview.Rating > 5 {
		http.Error(writer, "Fill all the fields properly.", http.StatusBadRequest)
		return
	}

	tourRatingPreview.ID = int(tourRatingPreview.ID)

	err = handler.TourRatingPreviewService.Create(&tourRatingPreview)
	if err != nil {
		log.Println("Error while creating new checkpoint:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
