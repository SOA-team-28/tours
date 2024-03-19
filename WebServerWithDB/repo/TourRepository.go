package repo

import (
	"database-example/model"
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