package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"

	"gorm.io/gorm"
)

type TourRatingPreviewService struct {
	TourRatingPreviewRepo *repo.TourRatingPreviewRepository
}

func NewTourRatingPreviewService(db *gorm.DB) *TourRatingPreviewService {
	tourRatingPreviewRepo := repo.NewTourRatingPreviewRepository(db)
	return &TourRatingPreviewService{TourRatingPreviewRepo: tourRatingPreviewRepo}
}

func (service *TourRatingPreviewService) Find(id int) (*model.TourRatingPreview, error) {
	tourRatingPreview, err := service.TourRatingPreviewRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("checkpoint with id %d not found", id)
	}
	return &tourRatingPreview, nil
}

func (service *TourRatingPreviewService) Create(tourRatingPreview *model.TourRatingPreview) error {
	err := service.TourRatingPreviewRepo.Create(tourRatingPreview)
	if err != nil {
		return err
	}
	return nil
}
