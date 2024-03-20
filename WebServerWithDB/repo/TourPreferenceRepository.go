package repo

import (
	"database-example/model"
	"gorm.io/gorm"
)

func NewTourPreferenceRepository(databaseConnection *gorm.DB) *TourPreferenceRepository {
	return &TourPreferenceRepository{DatabaseConnection: databaseConnection}
}

type TourPreferenceRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TourPreferenceRepository) FindAll() ([]model.TourPreference, error) {
	var tourPreferences []model.TourPreference
	dbResult := repo.DatabaseConnection.Find(&tourPreferences)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return tourPreferences, nil
}

func (repo *TourPreferenceRepository) Create(tourPreference *model.TourPreference) (*model.TourPreference, error) {
	dbResult := repo.DatabaseConnection.Create(tourPreference)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return tourPreference, nil
}
