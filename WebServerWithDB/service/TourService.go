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