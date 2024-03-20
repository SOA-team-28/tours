package service

import (
	"bytes"
	"database-example/model"
	"database-example/repo"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type MapObjectService struct {
	MapObjectRepo *repo.MapObjectRepository
}

func NewMapObjectService(db *gorm.DB) *MapObjectService {
	mapObjectRepo := repo.NewMapObjectRepository(db)
	return &MapObjectService{MapObjectRepo: mapObjectRepo}
}

func (service *MapObjectService) Find(id int) (*model.MapObject, error) {
	mapObject, err := service.MapObjectRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("equipment with id %d not found", id)
	}
	return &mapObject, nil
}

func (service *MapObjectService) FindAll() ([]model.MapObject, error) {
	return service.MapObjectRepo.FindAll()
}

func (service *MapObjectService) Create(mapObject *model.MapObject, userId int, status string) *model.MapObjectDTO {
	createdMapObject, _ := service.MapObjectRepo.Create(mapObject)
	createdMapObjectDTO := createdMapObject.MapToMapObjectDTO()

	mapObjectId := createdMapObject.ID

	if status == "Public" {
		// Kreiranje tela zahteva, ako je potrebno
		requestBody := []byte(`{"key": "value"}`) // Prilagodite telo zahteva vašim potrebama

		// Kreiranje HTTP zahteva
		url := fmt.Sprintf("https://localhost:44333/api/administration/objectRequests/create/%d/%d/%s", mapObjectId, userId, "OnHold")
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
		if err != nil {
			// Obrada greške pri kreiranju zahteva
			return nil
		}

		// Postavljanje zaglavlja zahteva, ako je potrebno
		req.Header.Set("Content-Type", "application/json")

		// Slanje zahteva
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			// Obrada greške pri slanju zahteva
			return nil
		}
		defer resp.Body.Close()
	}

	return &createdMapObjectDTO
}
