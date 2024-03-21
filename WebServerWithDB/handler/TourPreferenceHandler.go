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

type TourPreferenceHandler struct {
	TourPreferenceService *service.TourPreferenceService
}

func NewTourPreferenceHandler(db *gorm.DB) *TourPreferenceHandler {
	tourPreferenceService := service.NewTourPreferenceService(db)
	return &TourPreferenceHandler{TourPreferenceService: tourPreferenceService}
}

func (h *TourPreferenceHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/tourPreferences/get-by-creator/{id}", h.GetAllByCreatorId).Methods("GET")
	router.HandleFunc("/tourPreferences-get-all", h.GetAll).Methods("GET")
	router.HandleFunc("/tourPreferences", h.Create).Methods("POST")
	router.HandleFunc("/tourPreferences", h.Update).Methods("PUT")
	router.HandleFunc("/tourPreferences/{id}", h.Delete).Methods("DELETE")
}

func (handler *TourPreferenceHandler) GetAllByCreatorId(writer http.ResponseWriter, req *http.Request) {
	// Očitavanje parametara iz zahteva, ako je potrebno
	creatorIdString := mux.Vars(req)["id"]
	creatorId, err := strconv.Atoi(creatorIdString)
	if err != nil {
		log.Println("Error parsing creator ID:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// Poziv servisne metode za pronalaženje preferencija turista po ID-u kreatora
	tourPreferences, err := handler.TourPreferenceService.FindByCreatorId(creatorId)
	if err != nil {
		log.Println("Error retrieving TourPreferences:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Pretvaranje pronađenih turističkih preferencija u DTO objekte
	var tourPreferencesDTO []model.TourPreferenceDTO
	for _, pref := range tourPreferences {
		tourPreferencesDTO = append(tourPreferencesDTO, pref.MapToTourPreferenceDTO())
	}

	// Postavljanje zaglavlja odgovora na application/json
	writer.Header().Set("Content-Type", "application/json")
	// Slanje odgovora s pronađenim turističkim preferencijama u obliku JSON-a
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(tourPreferencesDTO)
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

func (handler *TourPreferenceHandler) Update(writer http.ResponseWriter, req *http.Request) {

	var tourPreferenceDTO model.TourPreferenceDTO
	err := json.NewDecoder(req.Body).Decode(&tourPreferenceDTO)
	if err != nil {
		log.Println("Error while parsing JSON:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	tourPreference := tourPreferenceDTO.MapToTourPreference()

	var updatedDTO = handler.TourPreferenceService.Update(&tourPreference)

	// Serijalizacija mapObjectDTO u JSON format
	jsonResponse, err := json.Marshal(updatedDTO)
	if err != nil {
		http.Error(writer, "Failed to serialize response", http.StatusInternalServerError)
		return
	}

	// Postavljanje Content-Type zaglavlja na application/json
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(jsonResponse)
}

func (handler *TourPreferenceHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	idString := params["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println("Error parsing ID:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.TourPreferenceService.Delete(id)
	if err != nil {
		log.Println("Error deleting TourPreference:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}
