package service

import (
	"database-example/model"
	"database-example/repo"
	"gorm.io/gorm"
	"log"
)

type TourPreferenceService struct {
	TourPreferenceRepo *repo.TourPreferenceRepository
}

func NewTourPreferenceService(db *gorm.DB) *TourPreferenceService {
	tourPreferenceRepo := repo.NewTourPreferenceRepository(db)
	return &TourPreferenceService{TourPreferenceRepo: tourPreferenceRepo}
}

func (service *TourPreferenceService) FindByCreatorId(id int) ([]model.TourPreference, error) {
	tourPreferences, err := service.TourPreferenceRepo.FindByCreatorId(id)
	if err != nil {
		// Obrada greške ako se dogodi prilikom poziva repozitorijuma
		return nil, err
	}
	return tourPreferences, nil
}

func (service *TourPreferenceService) FindAll() ([]model.TourPreference, error) {
	return service.TourPreferenceRepo.FindAll()
}

func (service *TourPreferenceService) Create(tourPreference *model.TourPreference) *model.TourPreferenceDTO {
	createdTourPreference, _ := service.TourPreferenceRepo.Create(tourPreference)
	createdTourPreferenceDTO := createdTourPreference.MapToTourPreferenceDTO()
	return &createdTourPreferenceDTO
}

func (service *TourPreferenceService) Update(tourPreference *model.TourPreference) *model.TourPreferenceDTO {
	updatedTourPreference, _ := service.TourPreferenceRepo.Update(tourPreference)
	updatedTourPreferenceDTO := updatedTourPreference.MapToTourPreferenceDTO()
	return &updatedTourPreferenceDTO
}

func (service *TourPreferenceService) Delete(id int) error {
	err := service.TourPreferenceRepo.Delete(id)
	if err != nil {
		log.Println("Error deleting TourPreference:", err)
		return err
	}
	return nil
}
