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

type CheckpointHandler struct {
	CheckpointService *service.CheckpointService
}

func NewCheckpointHandler(db *gorm.DB) *CheckpointHandler {
	checkpointService := service.NewCheckpointService(db)
	return &CheckpointHandler{CheckpointService: checkpointService}
}

func (h *CheckpointHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/checkpoints/{id}", h.Get).Methods("GET")
	router.HandleFunc("/checkpoints", h.Create).Methods("POST")
}

func (handler *CheckpointHandler) Get(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	idString := params["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println("Error parsing ID:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("Checkpoint with ID: %d", id)
	checkpoint, err := handler.CheckpointService.Find(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(checkpoint)
}

func (handler *CheckpointHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var checkpoint model.Checkpoint
	err := json.NewDecoder(req.Body).Decode(&checkpoint)
	if err != nil {
		log.Println("Error while parsing JSON:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	checkpoint.ID = int(checkpoint.ID)

	err = handler.CheckpointService.Create(&checkpoint)
	if err != nil {
		log.Println("Error while creating new checkpoint:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
