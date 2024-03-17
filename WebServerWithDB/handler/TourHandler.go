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
	var tourDTO model.TourDTO
	err := json.NewDecoder(req.Body).Decode(&tourDTO)
	if err != nil {
		log.Println("Error while parsing JSON:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	tour, equipment, checkpoints := tourDTO.MapToTour()

	// Now you have tour, equipment, and checkpoints to work with

		// Print all mapped fields
		log.Println("Tour:")
		log.Printf("ID: %d", tour.ID)
		log.Printf("Name: %s", tour.Name)
		log.Printf("Description: %s", tour.Description)
		log.Printf("DemandingLevel: %s", tour.DemandingLevel)
		log.Printf("Price: %d", tour.Price)
		log.Printf("Tags: %v", tour.Tags)
		log.Printf("AuthorID: %d", tour.AuthorID)
		log.Printf("Status: %s", tour.Status)
		log.Printf("TourTimes: %v", tour.TourTimes)
		log.Printf("TourRatings: %v", tour.TourRatings)
		log.Printf("Closed: %t", tour.Closed)
	
		log.Println("Equipment:")
		for _, e := range equipment {
			log.Printf("ID: %d, Name: %s, Description: %s", e.ID, e.Name, e.Description)
		}
	
		log.Println("Checkpoints:")
		for _, c := range checkpoints {
			log.Printf("ID: %d, TourID: %d, AuthorID: %d, Longitude: %f, Latitude: %f, Name: %s, Description: %s, Pictures: %s, RequiredTime: %d, CheckpointSecret: %s, EncounterID: %d, IsSecretPrerequisite: %t", c.ID, c.TourID, c.AuthorID, c.Longitude, c.Latitude, c.Name, c.Description, c.Pictures, c.RequiredTime, c.CheckpointSecret, c.EncounterID, c.IsSecretPrerequisite)
		}
		
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}


