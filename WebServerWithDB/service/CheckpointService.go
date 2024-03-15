package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"

	"gorm.io/gorm"
)

type CheckpointService struct {
	CheckpointRepo *repo.CheckpointRepository
}

func NewCheckpointService(db *gorm.DB) *CheckpointService {
	checkpointRepo := repo.NewCheckpointRepository(db)
	return &CheckpointService{CheckpointRepo: checkpointRepo}
}

func (service *CheckpointService) Find(id int) (*model.Checkpoint, error) {
	checkpoint, err := service.CheckpointRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("checkpoint with id %d not found", id)
	}
	return &checkpoint, nil
}

func (service *CheckpointService) Create(checkpoint *model.Checkpoint) error {
	err := service.CheckpointRepo.Create(checkpoint)
	if err != nil {
		return err
	}
	return nil
}
