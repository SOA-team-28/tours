package repo

import (
	"database-example/model"
	"errors"

	"gorm.io/gorm"
)

func NewTourRepository(databaseConnection *gorm.DB) *TourRepository {
	return &TourRepository{DatabaseConnection: databaseConnection}
}

type TourRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TourRepository) FindById(id int) (model.Tour, error) {
	tour := model.Tour{}
	dbResult := repo.DatabaseConnection.Preload("Equipment").Preload("Checkpoints").First(&tour, "id = ?", id)
	if dbResult.Error != nil {
		return tour, dbResult.Error
	}
	return tour, nil
}

func (repo *TourRepository) Create(tour *model.Tour) error {
	dbResult := repo.DatabaseConnection.Create(tour)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *TourRepository) Update(tour *model.Tour) error {
	dbResult := repo.DatabaseConnection.Save(tour)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

func (repo *TourRepository) Delete(id int) error {
	dbResult := repo.DatabaseConnection.Delete(&model.Tour{}, id)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *TourRepository) GetToursByAuthor(authorID int) ([]model.Tour, error) {
    var tours []model.Tour
    // Assuming your Tour model has an AuthorID field
    result := repo.DatabaseConnection.Where("author_id = ?", authorID).Find(&tours)
    if result.Error != nil {
        return nil, result.Error
    }
    return tours, nil
}

func (repo *TourRepository) GetAll(page, pageSize int) ([]model.Tour, error) {
	var tours []model.Tour
	dbResult := repo.DatabaseConnection.Limit(pageSize).Offset((page - 1) * pageSize).
		Preload("Equipment").Preload("Checkpoints").
		Find(&tours)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return tours, nil
}

func (repo *TourRepository) AddEquipment(tourID, equipmentID int) error {
	var tour model.Tour
	dbResult := repo.DatabaseConnection.Preload("Equipment").First(&tour, "id = ?", tourID)
	if dbResult.Error != nil {
		return dbResult.Error
	}

	var equipment model.Equipment
	dbResult = repo.DatabaseConnection.First(&equipment, "id = ?", equipmentID)
	if dbResult.Error != nil {
		return dbResult.Error
	}

	tour.Equipment = append(tour.Equipment, equipment)
	dbResult = repo.DatabaseConnection.Save(&tour)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

func (repo *TourRepository) RemoveEquipment(tourID, equipmentID int) error {
	// Find the tour
	var tour model.Tour
	if err := repo.DatabaseConnection.Preload("Equipment").First(&tour, "id = ?", tourID).Error; err != nil {
		return err
	}

	// Check if the equipment exists in the tour
	var updatedEquipment []model.Equipment
	var found bool
	for _, eq := range tour.Equipment {
		if eq.ID == equipmentID {
			found = true
		} else {
			updatedEquipment = append(updatedEquipment, eq)
		}
	}

	// If equipment not found in the tour, return an error
	if !found {
		return errors.New("equipment not found in the tour")
	}

	// Update the equipment list
	tour.Equipment = updatedEquipment

	// Save the updated tour
	if err := repo.DatabaseConnection.Save(&tour).Error; err != nil {
		return err
	}

	// Delete the corresponding tour_equipments row
	if err := repo.DatabaseConnection.Exec("DELETE FROM tour_equipments WHERE tour_id = ? AND equipment_id = ?", tourID, equipmentID).Error; err != nil {
		return err
	}

	return nil
}