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

	var tourRatingPreviewDTO model.TourRatingPreviewDTO
	err := json.NewDecoder(req.Body).Decode(&tourRatingPreviewDTO)
	if err != nil {
		log.Println("Error while parsing JSON:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if tourRatingPreviewDTO.TourID == 0 || tourRatingPreviewDTO.TouristID == 0 || tourRatingPreviewDTO.Rating == 0 || tourRatingPreviewDTO.Rating > 5 {
		http.Error(writer, "Fill all the fields properly.", http.StatusBadRequest)
		return
	}

	// sprovesti sledece validacije kad to bude moguce
	//_tourOwnershipService.GetPurchasedToursByUser
	//_executionRepository.GetExactExecution - IsTourProgressAbove35Percent(), HasOneWeekPassedSinceLastActivity()

	tourService := &service.TourService{}
	tour, _ := tourService.Find(tourRatingPreviewDTO.TourID)
	var tourModel model.Tour
	if tour != nil {
		tourModel = *tour
	}

	tourRatingPreview := tourRatingPreviewDTO.MapToTourRatingPreview(tourModel)

	err = handler.TourRatingPreviewService.Create(&tourRatingPreview)
	if err != nil {
		log.Println("Error while creating new tour rating preview:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	responseBody, err := json.Marshal(tourRatingPreview)
	if err != nil {
		log.Println("Error marshaling TourRatingPreview:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(responseBody)
}
