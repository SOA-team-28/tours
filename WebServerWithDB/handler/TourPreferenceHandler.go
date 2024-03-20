package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type TourPreferenceHandler struct {
	TourPreferenceService *service.TourPreferenceService
}

func NewTourPreferenceHandler(db *gorm.DB) *TourPreferenceHandler {
	tourPreferenceService := service.NewTourPreferenceService(db)
	return &TourPreferenceHandler{TourPreferenceService: tourPreferenceService}
}

func (h *TourPreferenceHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/tourPreferences-get-all", h.GetAll).Methods("GET")
	router.HandleFunc("/tourPreferences", h.Create).Methods("POST")
}

func (handler *TourPreferenceHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	tourPreferences, err := handler.TourPreferenceService.FindAll()
	if err != nil {
		log.Println("Error retrieving MapObjects:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Pretvaranje mapObjects u DTO objekte
	var tourPreferencesDTO []model.TourPreferenceDTO
	for _, obj := range tourPreferences {
		tourPreferencesDTO = append(tourPreferencesDTO, obj.MapToTourPreferenceDTO())
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(tourPreferencesDTO)
}

func (handler *TourPreferenceHandler) Create(writer http.ResponseWriter, req *http.Request) {

	var tourPreferenceDTO model.TourPreferenceDTO
	err := json.NewDecoder(req.Body).Decode(&tourPreferenceDTO)
	if err != nil {
		log.Println("Error while parsing JSON:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	/*
		params := mux.Vars(req)
		status := params["status"]
		userIdString := params["userId"]
		userId, err := strconv.Atoi(userIdString)
		if err != nil {
			log.Println("Error parsing ID:", err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
	*/
	tourPreference := tourPreferenceDTO.MapToTourPreference()

	var createdDTO = handler.TourPreferenceService.Create(&tourPreference)

	// Serijalizacija mapObjectDTO u JSON format
	jsonResponse, err := json.Marshal(createdDTO)
	if err != nil {
		http.Error(writer, "Failed to serialize response", http.StatusInternalServerError)
		return
	}

	// Postavljanje Content-Type zaglavlja na application/json
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(jsonResponse)
}
