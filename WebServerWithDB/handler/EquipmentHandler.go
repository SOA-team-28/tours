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

type EquipmentHandler struct {
    EquipmentService *service.EquipmentService
}

func NewEquipmentHandler(db *gorm.DB) *EquipmentHandler {
    equipmentService := service.NewEquipmentService(db)
    return &EquipmentHandler{EquipmentService: equipmentService}
}

func (h *EquipmentHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/equipments/{id}", h.Get).Methods("GET")
	router.HandleFunc("/equipments", h.Create).Methods("POST")
	router.HandleFunc("/equipments/available/{tourId}", h.GetAvailableEquipment).Methods("POST")
}

func (handler *EquipmentHandler) Get(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	idString := params["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println("Error parsing ID:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("Equipment with ID: %d", id)
	equipment, err := handler.EquipmentService.Find(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(equipment)
}

func (handler *EquipmentHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var equipment model.Equipment
	err := json.NewDecoder(req.Body).Decode(&equipment)
	if err != nil {
		log.Println("Error while parsing JSON:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	equipment.ID = int(equipment.ID)

	err = handler.EquipmentService.Create(&equipment)
	if err != nil {
		log.Println("Error while creating new equipment:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *EquipmentHandler) GetAvailableEquipment(writer http.ResponseWriter, req *http.Request) {
	// Parse tour ID from URL parameters
	params := mux.Vars(req)
	tourIDString := params["tourId"]
	tourID, err := strconv.Atoi(tourIDString)
	if err != nil {
		log.Println("Error parsing tour ID:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var equipmentIDs []int
	err = json.NewDecoder(req.Body).Decode(&equipmentIDs)
	if err != nil {
		log.Println("Error decoding equipment IDs:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	availableEquipment, err := handler.EquipmentService.GetAvailableEquipment(tourID, equipmentIDs)
	if err != nil {
		log.Println("Error retrieving available equipment:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(availableEquipment)
}