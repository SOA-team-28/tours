package model

import (
	"database-example/service"

	"strings"
	"time"
)

type CheckpointDTO struct {
	ID                   int      `json:"id"`
	TourID               int      `json:"tourId"`
	AuthorID             int      `json:"authorId"`
	Longitude            float64  `json:"longitude"`
	Latitude             float64  `json:"latitude"`
	Name                 string   `json:"name"`
	Description          string   `json:"description"`
	Pictures             []string `json:"pictures"`
	RequiredTime         int      `json:"requiredTimeInSeconds"`
	CheckpointSecret     string   `json:"checkpointSecret"`
	EncounterID          int      `json:"encounterId"`
	IsSecretPrerequisite bool     `json:"isSecretPrerequisite"`
}

// TourData represents the JSON data structure for creating a tour
type TourDTO struct {
	ID             int             `json:"id"`
	Name           string          `json:"name"`
	Description    string          `json:"description"`
	DemandingLevel string          `json:"demandignessLevel"`
	Price          int             `json:"price"`
	Tags           []string        `json:"tags"`
	AuthorID       int             `json:"authorId"`
	Status         string          `json:"status"`
	Equipment      []Equipment     `json:"equipment"`
	Checkpoints    []CheckpointDTO `json:"checkpoints"`
	TourTimes      []int           `json:"tourTimes"`
	TourRatings    []int           `json:"tourRatings"`
	Closed         bool            `json:"closed"`
}

type TourRatingPreviewDTO struct {
	ID           int       `json:"id"`
	Rating       int       `json:"rating"`
	Comment      string    `json:"comment"`
	TouristID    int       `json:"touristId"`
	TourID       int       `gorm:"tourId"`
	TourDate     time.Time `json:"tourDate"`
	CreationDate time.Time `json:"creationDate"`
	ImageNames   []string  `json:"imageNames"`
}

func (tourRatingPreviewDTO *TourRatingPreviewDTO) MapToTourRatingPreview() TourRatingPreview {

	tourService := &service.TourService{}
	tour, _ := tourService.Find(tourRatingPreviewDTO.TourID)
	var tourModel Tour
	if tour != nil {
		tourModel = *tour
	}

	tourRatingPreview := TourRatingPreview{
		ID:           tourRatingPreviewDTO.ID,
		Rating:       tourRatingPreviewDTO.Rating,
		Comment:      tourRatingPreviewDTO.Comment,
		TouristID:    tourRatingPreviewDTO.TouristID,
		Tour:         tourModel,
		TourDate:     tourRatingPreviewDTO.TourDate,
		CreationDate: tourRatingPreviewDTO.CreationDate,
		ImageNames:   strings.Join(tourRatingPreviewDTO.ImageNames, "|"),
	}

	return tourRatingPreview
}

func (tourDTO *TourDTO) MapToTour() (Tour, []Equipment, []Checkpoint) {

	var equipment []Equipment
	for _, eq := range tourDTO.Equipment {
		equipment = append(equipment, Equipment{
			ID:          eq.ID,
			Name:        eq.Name,
			Description: eq.Description,
		})
	}

	var checkpoints []Checkpoint
	for _, cp := range tourDTO.Checkpoints {
		checkpoints = append(checkpoints, Checkpoint{
			ID:                   cp.ID,
			TourID:               cp.TourID,
			AuthorID:             cp.AuthorID,
			Longitude:            cp.Longitude,
			Latitude:             cp.Latitude,
			Name:                 cp.Name,
			Description:          cp.Description,
			Pictures:             strings.Join(cp.Pictures, "|"), // Join pictures into a single string
			RequiredTime:         cp.RequiredTime,
			CheckpointSecret:     cp.CheckpointSecret,
			EncounterID:          cp.EncounterID,
			IsSecretPrerequisite: cp.IsSecretPrerequisite,
		})
	}

	tour := Tour{
		ID:             tourDTO.ID,
		Name:           tourDTO.Name,
		Description:    tourDTO.Description,
		DemandingLevel: tourDTO.DemandingLevel,
		Price:          tourDTO.Price,
		Tags:           strings.Join(tourDTO.Tags, "|"),
		AuthorID:       tourDTO.AuthorID,
		Status:         tourDTO.Status,
		Checkpoints:    checkpoints,
		Equipment:      equipment,
		TourTimes:      "TODO",
		TourRatings:    "TODO",
		Closed:         tourDTO.Closed,
	}

	return tour, equipment, checkpoints
}
