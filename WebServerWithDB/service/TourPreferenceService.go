package service

import (
	"database-example/model"
	"database-example/repo"
	"gorm.io/gorm"
)

type TourPreferenceService struct {
	TourPreferenceRepo *repo.TourPreferenceRepository
}

func NewTourPreferenceService(db *gorm.DB) *TourPreferenceService {
	tourPreferenceRepo := repo.NewTourPreferenceRepository(db)
	return &TourPreferenceService{TourPreferenceRepo: tourPreferenceRepo}
}

func (service *TourPreferenceService) FindAll() ([]model.TourPreference, error) {
	return service.TourPreferenceRepo.FindAll()
}

func (service *TourPreferenceService) Create(tourPreference *model.TourPreference) *model.TourPreferenceDTO {
	createdTourPreference, _ := service.TourPreferenceRepo.Create(tourPreference)
	createdTourPreferenceDTO := createdTourPreference.MapToTourPreferenceDTO()
	return &createdTourPreferenceDTO
}
