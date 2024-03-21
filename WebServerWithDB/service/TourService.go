package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"

	"gorm.io/gorm"
)

type TourService struct {
	TourRepo *repo.TourRepository
}

func NewTourService(db *gorm.DB) *TourService {
	tourRepo := repo.NewTourRepository(db)
	return &TourService{TourRepo: tourRepo}
}

func (service *TourService) Find(id int) (*model.Tour, error) {
	tour, err := service.TourRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("tour with id %d not found", id)
	}
	return &tour, nil
}

func (service *TourService) Create(tour *model.Tour) error {
	err := service.TourRepo.Create(tour)
	if err != nil {
		return err
	}
	return nil
}

func (service *TourService) Update(tour *model.Tour) error {
	err := service.TourRepo.Update(tour)
	if err != nil {
		return err
	}
	return nil
}

func (service *TourService) Delete(id int) error {
	err := service.TourRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (service *TourService) GetToursByAuthor(authorID, page, pageSize int) ([]model.Tour, error) {
	// Implement logic to get tours by author ID from the repository
	tours, err := service.TourRepo.GetToursByAuthor(authorID)
	if err != nil {
		return nil, err
	}
	return tours, nil
}

func (service *TourService) GetAll(page, pageSize int) ([]model.Tour, error) {
	// Implement logic to get all tours from the repository
	tours, err := service.TourRepo.GetAll(page, pageSize)
	if err != nil {
		return nil, err
	}
	return tours, nil
}

func (service *TourService) AddEquipment(tourID, equipmentID int) error {
	// Implement logic to add equipment to a tour in the repository
	err := service.TourRepo.AddEquipment(tourID, equipmentID)
	if err != nil {
		return err
	}
	return nil
}

func (service *TourService) RemoveEquipment(tourID, equipmentID int) error {
	// Implement logic to remove equipment from a tour in the repository
	err := service.TourRepo.RemoveEquipment(tourID, equipmentID)
	if err != nil {
		return err
	}
	return nil
}
