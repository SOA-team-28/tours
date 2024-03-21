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
	router.HandleFunc("/tours/{id}", h.Update).Methods("PUT")
	router.HandleFunc("/tours/{id}", h.Delete).Methods("DELETE")
	router.HandleFunc("/tours/byauthor/{authorId}", h.GetToursByAuthor).Methods("GET")
	router.HandleFunc("/tours", h.GetAll).Methods("GET")
	router.HandleFunc("/tours/add/{tourId}/{equipmentId}", h.AddEquipment).Methods("PUT")
	router.HandleFunc("/tours/remove/{tourId}/{equipmentId}", h.RemoveEquipment).Methods("PUT")
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
	tourDTO := model.TourToDTO(tour)
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(tourDTO)
}

func (handler *TourHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var tourDTO model.TourDTO
	err := json.NewDecoder(req.Body).Decode(&tourDTO)
	if err != nil {
		log.Println("Error while parsing JSON:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	tour, _, _ := tourDTO.MapToTour()

	err = handler.TourService.Create(&tour)
	if err != nil {
		log.Println("Error while creating new tour:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	tourDTO = model.TourToDTO(&tour)
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(tourDTO)

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *TourHandler) Update(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	idString := params["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println("Error parsing ID:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var tourDTO model.TourDTO
	err = json.NewDecoder(req.Body).Decode(&tourDTO)
	if err != nil {
		log.Println("Error while parsing JSON:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	tour, _, _ := tourDTO.MapToTour()
	tour.ID = id

	err = handler.TourService.Update(&tour)
	if err != nil {
		log.Println("Error while updating tour:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(tourDTO)
}

func (handler *TourHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	idString := params["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println("Error parsing ID:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.TourService.Delete(id)
	if err != nil {
		log.Println("Error while deleting tour:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}

func (handler *TourHandler) GetAll(writer http.ResponseWriter, req *http.Request) {

	// Implement logic to get all tours from the service layer
	tours, err := handler.TourService.GetAll(0, 0)
	if err != nil {
		// Handle error
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write response
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(tours)
}

func (handler *TourHandler) GetToursByAuthor(writer http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	idString := params["authorId"]
	authorId, err := strconv.Atoi(idString)
	if err != nil {
		log.Println("Error parsing ID:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// Implement logic to get tours by author from the service layer
	tours, err := handler.TourService.GetToursByAuthor(authorId, 0, 0)
	if err != nil {
		// Handle error
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	var tourDTOs []model.TourDTO
	for _, tour := range tours {
		tourDTO := model.TourToDTO(&tour)
		tourDTOs = append(tourDTOs, tourDTO)
	}

	// Write response
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(tourDTOs)
}

func (handler *TourHandler) AddEquipment(writer http.ResponseWriter, req *http.Request) {
	// Parse tour ID and equipment ID from URL parameters
	params := mux.Vars(req)
	tourID, err := strconv.Atoi(params["tourId"])
	if err != nil {
		// Handle error
		return
	}
	equipmentID, err := strconv.Atoi(params["equipmentId"])
	if err != nil {
		// Handle error
		return
	}

	// Implement logic to add equipment to the tour in the service layer
	err = handler.TourService.AddEquipment(tourID, equipmentID)
	if err != nil {
		// Handle error
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write response
	writer.WriteHeader(http.StatusNoContent)
}

func (handler *TourHandler) RemoveEquipment(writer http.ResponseWriter, req *http.Request) {
	// Parse tour ID and equipment ID from URL parameters
	params := mux.Vars(req)
	tourID, err := strconv.Atoi(params["tourId"])
	if err != nil {
		// Handle error
		return
	}
	equipmentID, err := strconv.Atoi(params["equipmentId"])
	if err != nil {
		// Handle error
		return
	}

	// Implement logic to remove equipment from the tour in the service layer
	err = handler.TourService.RemoveEquipment(tourID, equipmentID)
	if err != nil {
		// Handle error
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write response
	writer.WriteHeader(http.StatusNoContent)
}
