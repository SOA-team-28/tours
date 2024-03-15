package repo

import (
	"database-example/model"
	"gorm.io/gorm"
)

func NewCheckpointRepository(databaseConnection *gorm.DB) *CheckpointRepository {
	return &CheckpointRepository{DatabaseConnection: databaseConnection}
}

type CheckpointRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *CheckpointRepository) FindById(id int) (model.Checkpoint, error) {
	checkpoint := model.Checkpoint{}
	dbResult := repo.DatabaseConnection.First(&checkpoint, "id = ?", id)
	if dbResult.Error != nil {
		return checkpoint, dbResult.Error
	}
	return checkpoint, nil
}

func (repo *CheckpointRepository) Create(checkpoint *model.Checkpoint) error {
	dbResult := repo.DatabaseConnection.Create(checkpoint)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
