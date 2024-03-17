package repo

import (
	"database-example/model"
	"gorm.io/gorm"
)

func NewTourRatingPreviewRepository(databaseConnection *gorm.DB) *TourRatingPreviewRepository {
	return &TourRatingPreviewRepository{DatabaseConnection: databaseConnection}
}

type TourRatingPreviewRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TourRatingPreviewRepository) FindById(id int) (model.TourRatingPreview, error) {
	tourRatingPreview := model.TourRatingPreview{}
	dbResult := repo.DatabaseConnection.First(&tourRatingPreview, "id = ?", id)
	if dbResult.Error != nil {
		return tourRatingPreview, dbResult.Error
	}
	return tourRatingPreview, nil
}

func (repo *TourRatingPreviewRepository) Create(tourRatingPreview *model.TourRatingPreview) error {
	dbResult := repo.DatabaseConnection.Create(tourRatingPreview)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
