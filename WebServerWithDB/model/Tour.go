package model

import "strings"

type MapObjectType int

const (
	Other MapObjectType = iota
	Restaurant
	WC
	Parking
)

type Tour struct {
	ID             int          `gorm:"primaryKey" json:"id"`
	Name           string       `json:"name"`
	Description    string       `json:"description"`
	DemandingLevel string       `json:"demandignessLevel"`
	Price          float64      `json:"price"`
	Tags           string       `json:"tags"`
	AuthorID       int          `json:"authorId"`
	Status         string       `json:"status"`
	Equipment      []Equipment  `gorm:"many2many:tour_equipments;" json:"equipment"`
	Checkpoints    []Checkpoint `gorm:"many2many:tour_checkpoints;" json:"checkpoints"`
	TourTimes      string       `json:"tourTimes"`
	TourRatings    string       `json:"tourRatings"`
	Closed         bool         `json:"closed"`
}

type Checkpoint struct {
	ID                   int     `gorm:"primaryKey" json:"id"`
	TourID               int     `json:"tourId"`
	AuthorID             int     `json:"authorId"`
	Longitude            float64 `json:"longitude"`
	Latitude             float64 `json:"latitude"`
	Name                 string  `json:"name"`
	Description          string  `json:"description"`
	Pictures             string  `json:"pictures"`
	RequiredTime         float64 `json:"requiredTimeInSeconds"`
	CheckpointSecret     string  `json:"checkpointSecret"`
	EncounterID          int     `json:"encounterId"`
	IsSecretPrerequisite bool    `json:"isSecretPrerequisite"`
}

type MapObject struct {
	ID          int           `gorm:"primaryKey" json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	PictureURL  string        `json:"pictureURL"`
	Category    MapObjectType `json:"category"`
	Longitude   float64       `json:"longitude"`
	Latitude    float64       `json:"latitude"`
}

type Equipment struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type TourCheckpoint struct {
	TourID       int `json:"tourId"`
	CheckpointID int `json:"checkpointId"`
}

type TourEquipment struct {
	TourID      int `json:"tourId"`
	EquipmentID int `json:"equipmentId"`
}

func TourToDTO(tour *Tour) TourDTO {
	var tags []string
	if tour.Tags != "" {
		tags = strings.Split(tour.Tags, "|")
	}

	var equipmentDTOs []Equipment
	for _, eq := range tour.Equipment {
		equipmentDTOs = append(equipmentDTOs, Equipment{
			ID:          eq.ID,
			Name:        eq.Name,
			Description: eq.Description,
		})
	}

	var checkpointDTOs []CheckpointDTO
	for _, cp := range tour.Checkpoints {
		pictures := strings.Split(cp.Pictures, "|")
		checkpointDTOs = append(checkpointDTOs, CheckpointDTO{
			ID:                   cp.ID,
			TourID:               cp.TourID,
			AuthorID:             cp.AuthorID,
			Longitude:            cp.Longitude,
			Latitude:             cp.Latitude,
			Name:                 cp.Name,
			Description:          cp.Description,
			Pictures:             pictures,
			RequiredTime:         cp.RequiredTime,
			CheckpointSecret:     cp.CheckpointSecret,
			EncounterID:          cp.EncounterID,
			IsSecretPrerequisite: cp.IsSecretPrerequisite,
		})
	}

	// Initialize Equipment and Checkpoints as empty slices if they are nil
	if equipmentDTOs == nil {
		equipmentDTOs = make([]Equipment, 0)
	}
	if checkpointDTOs == nil {
		checkpointDTOs = make([]CheckpointDTO, 0)
	}

	return TourDTO{
		ID:             tour.ID,
		Name:           tour.Name,
		Description:    tour.Description,
		DemandingLevel: tour.DemandingLevel,
		Price:          tour.Price,
		Tags:           tags,
		AuthorID:       tour.AuthorID,
		Status:         tour.Status,
		Equipment:      equipmentDTOs,
		Checkpoints:    checkpointDTOs,
		TourTimes:      nil,
		TourRatings:    nil,
		Closed:         tour.Closed,
	}
}

func (mapObject *MapObject) MapToMapObjectDTO() MapObjectDTO {
	var category string

	switch mapObject.Category {
	case Other:
		category = "Other"
	case Restaurant:
		category = "Restaurant"
	case WC:
		category = "WC"
	default:
		category = "Parking"
	}

	mapObjectDTO := MapObjectDTO{
		ID:          mapObject.ID,
		Name:        mapObject.Name,
		Description: mapObject.Description,
		PictureURL:  mapObject.PictureURL,
		Category:    category,
		Longitude:   mapObject.Longitude,
		Latitude:    mapObject.Latitude,
	}

	return mapObjectDTO
}
