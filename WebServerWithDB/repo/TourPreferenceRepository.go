package repo

import (
	"database-example/model"
	"errors"
	"gorm.io/gorm"
)

func NewTourPreferenceRepository(databaseConnection *gorm.DB) *TourPreferenceRepository {
	return &TourPreferenceRepository{DatabaseConnection: databaseConnection}
}

type TourPreferenceRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TourPreferenceRepository) FindByCreatorId(id int) ([]model.TourPreference, error) {
	var tourPreferences []model.TourPreference
	dbResult := repo.DatabaseConnection.Find(&tourPreferences, "creator_id = ?", id)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return tourPreferences, nil
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
	validateTransportation(tourPreference.Walk, tourPreference.Bike, tourPreference.Car, tourPreference.Boat)
	dbResult := repo.DatabaseConnection.Create(tourPreference)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return tourPreference, nil
}

func validateTransportation(walk, bike, car, boat int) {
	if walk < 0 || walk > 3 {
		panic(errors.New("Invalid value for Walk. It should be between 0 and 3."))
	}
	if bike < 0 || bike > 3 {
		panic(errors.New("Invalid value for Bike. It should be between 0 and 3."))
	}
	if car < 0 || car > 3 {
		panic(errors.New("Invalid value for Car. It should be between 0 and 3."))
	}
	if boat < 0 || boat > 3 {
		panic(errors.New("Invalid value for Boat. It should be between 0 and 3."))
	}
}
