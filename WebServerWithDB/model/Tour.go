package model

import "time"

type ExecutionStatus int

const (
	Completed ExecutionStatus = iota
	Abandoned
	InProgress
)

type Tour struct {
	ID             int          `gorm:"primaryKey" json:"id"`
	Name           string       `json:"name"`
	Description    string       `json:"description"`
	DemandingLevel string       `json:"demandignessLevel"`
	Price          int          `json:"price"`
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
	RequiredTime         int     `json:"requiredTimeInSeconds"`
	CheckpointSecret     string  `json:"checkpointSecret"`
	EncounterID          int     `json:"encounterId"`
	IsSecretPrerequisite bool    `json:"isSecretPrerequisite"`
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

type TourRatingPreview struct {
	ID           int       `gorm:"primaryKey" json:"id"`
	Rating       int       `json:"rating"`
	Comment      string    `json:"comment"`
	TourDate     time.Time `json:"tourDate"`
	CreationDate time.Time `json:"creationDate"`
	ImageNames   string    `json:"imageNames"`
	TouristID    int       `json:"touristId"`
	Tour         Tour      `gorm:"foreignKey:Tour"`
}
