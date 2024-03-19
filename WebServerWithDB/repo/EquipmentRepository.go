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

func (repo *EquipmentRepository) Create(equipment *model.Equipment) error {
	dbResult := repo.DatabaseConnection.Create(equipment)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *EquipmentRepository) FindAvailableByTourID(tourID int, excludedEquipmentIDs []int) ([]model.Equipment, error) {
    var availableEquipment []model.Equipment
	db := repo.DatabaseConnection

	// Retrieve all equipment IDs associated with the tour
	var tourEquipment []model.TourEquipment
	if err := db.Where("tour_id = ?", tourID).Find(&tourEquipment).Error; err != nil {
		return nil, err
	}

	// Store the equipment IDs associated with the tour
	var tourEquipmentIDs []int
	for _, te := range tourEquipment {
		tourEquipmentIDs = append(tourEquipmentIDs, te.EquipmentID)
	}


	// Check if there are no equipment IDs associated with the tour
	if len(tourEquipmentIDs) == 0 {
		// Query all equipment from the equipment table
		if err := db.Find(&availableEquipment).Error; err != nil {
			return nil, err
		}
	} else {
		// Query the equipment table to find equipment not in the tourEquipmentIDs list
		if err := db.Not("id IN ?", tourEquipmentIDs).Find(&availableEquipment).Error; err != nil {
			return nil, err
		}
	}

	return availableEquipment, nil
}