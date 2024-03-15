package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"

	"gorm.io/gorm"
)

type EquipmentService struct {
    EquipmentRepo *repo.EquipmentRepository
}


func NewEquipmentService(db *gorm.DB) *EquipmentService {
    equipmentRepo := repo.NewEquipmentRepository(db) 
    return &EquipmentService{EquipmentRepo: equipmentRepo}
}

func (service *EquipmentService) FindEquipment(id int) (*model.Equipment, error) {
	equipment, err := service.EquipmentRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("equipment with id %s not found", id)
	}
	return &equipment, nil
}

func (service *EquipmentService) Create(equipment *model.Equipment) error {
	err := service.EquipmentRepo.CreateEquipment(equipment)
	if err != nil {
		return err
	}
	return nil
}
