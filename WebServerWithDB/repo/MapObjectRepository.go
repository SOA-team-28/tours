package repo

import (
	"database-example/model"
	"gorm.io/gorm"
)

func NewMapObjectRepository(databaseConnection *gorm.DB) *MapObjectRepository {
	return &MapObjectRepository{DatabaseConnection: databaseConnection}
}

type MapObjectRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *MapObjectRepository) FindById(id int) (model.MapObject, error) {
	mapObject := model.MapObject{}
	dbResult := repo.DatabaseConnection.First(&mapObject, "id = ?", id)
	if dbResult.Error != nil {
		return mapObject, dbResult.Error
	}
	return mapObject, nil
}

func (repo *MapObjectRepository) FindAll() ([]model.MapObject, error) {
	var mapObjects []model.MapObject
	dbResult := repo.DatabaseConnection.Find(&mapObjects)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return mapObjects, nil
}

func (repo *MapObjectRepository) Create(mapObject *model.MapObject) (*model.MapObject, error) {
	dbResult := repo.DatabaseConnection.Create(mapObject)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return mapObject, nil
}
