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

type MapObjectHandler struct {
	MapObjectService *service.MapObjectService
}

func NewMapObjectHandler(db *gorm.DB) *MapObjectHandler {
	mapObjectService := service.NewMapObjectService(db)
	return &MapObjectHandler{MapObjectService: mapObjectService}
}

func (h *MapObjectHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/mapObjects/{id}", h.Get).Methods("GET")
	router.HandleFunc("/mapObjects-get-all", h.GetAll).Methods("GET")
	router.HandleFunc("/mapObjects/{userId}/{status}", h.Create).Methods("POST")
}

func (handler *MapObjectHandler) Get(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	idString := params["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println("Error parsing ID:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("Equipment with ID: %d", id)
	mapObject, err := handler.MapObjectService.Find(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(mapObject)
}

func (handler *MapObjectHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	mapObjects, err := handler.MapObjectService.FindAll()
	if err != nil {
		log.Println("Error retrieving MapObjects:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Pretvaranje mapObjects u DTO objekte
	var mapObjectsDTO []model.MapObjectDTO
	for _, obj := range mapObjects {
		mapObjectsDTO = append(mapObjectsDTO, obj.MapToMapObjectDTO())
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(mapObjectsDTO)
}

func (handler *MapObjectHandler) Create(writer http.ResponseWriter, req *http.Request) {
	/*
		// Dobijanje JWT tokena iz zaglavlja zahteva
		token, e := GetJWTToken(req)
		if e != nil {
			// Obrada greške ako token nije pronađen
			http.Error(writer, "Unauthorized", http.StatusUnauthorized)
			return
		}
	*/
	var mapObjectDTO model.MapObjectDTO
	err := json.NewDecoder(req.Body).Decode(&mapObjectDTO)
	if err != nil {
		log.Println("Error while parsing JSON:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	params := mux.Vars(req)
	status := params["status"]
	userIdString := params["userId"]
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		log.Println("Error parsing ID:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	mapObject := mapObjectDTO.MapToMapObject()

	var createdDTO = handler.MapObjectService.Create(&mapObject, userId, status /*, token*/)

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

/*
func GetJWTToken(req *http.Request) (string, error) {
	// Dobijanje vrednosti Authorization zaglavlja iz zahteva
	authHeader := req.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("Authorization header is missing")
	}

	// Provera da li je Authorization zaglavlje u formatu "Bearer <token>"
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("Invalid Authorization header format")
	}

	// Vraćanje tokena
	return parts[1], nil
}
*/
