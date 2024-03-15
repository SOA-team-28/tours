package repo

import (
	"database-example/model"
	"gorm.io/gorm"
)

func NewEquipmentRepository(databaseConnection *gorm.DB) *EquipmentRepository {
	return &EquipmentRepository{DatabaseConnection: databaseConnection}
}

type EquipmentRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *EquipmentRepository) FindById(id int) (model.Equipment, error) {
	equipment := model.Equipment{}
	dbResult := repo.DatabaseConnection.First(&equipment, "id = ?", id)
	if dbResult.Error != nil {
		return equipment, dbResult.Error
	}
	return equipment, nil
}

func (repo *EquipmentRepository) CreateEquipment(equipment *model.Equipment) error {
	dbResult := repo.DatabaseConnection.Create(equipment)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
