package model

import "strings"

type CheckpointDTO struct {
	ID                   int     `json:"id"`
	TourID               int     `json:"tourId"`
	AuthorID             int     `json:"authorId"`
	Longitude            float64 `json:"longitude"`
	Latitude             float64 `json:"latitude"`
	Name                 string  `json:"name"`
	Description          string  `json:"description"`
	Pictures             []string  `json:"pictures"`
	RequiredTime         int     `json:"requiredTimeInSeconds"`
	CheckpointSecret     string  `json:"checkpointSecret"`
	EncounterID          int     `json:"encounterId"`
	IsSecretPrerequisite bool    `json:"isSecretPrerequisite"`
}

// TourData represents the JSON data structure for creating a tour
type TourDTO struct {
	ID             int            `json:"id"`
	Name           string         `json:"name"`
	Description    string         `json:"description"`
	DemandingLevel string         `json:"demandignessLevel"`
	Price          int            `json:"price"`
	Tags           []string       `json:"tags"`
	AuthorID       int            `json:"authorId"`
	Status         string         `json:"status"`
	Equipment      []Equipment `json:"equipment"`
	Checkpoints    []CheckpointDTO `json:"checkpoints"`
	TourTimes      []int          `json:"tourTimes"`
	TourRatings    []int          `json:"tourRatings"`
	Closed         bool           `json:"closed"`
}

func (tourDTO *TourDTO) MapToTour() (Tour, []Equipment, []Checkpoint) {
	tour := Tour{
		ID:             tourDTO.ID,
		Name:           tourDTO.Name,
		Description:    tourDTO.Description,
		DemandingLevel: tourDTO.DemandingLevel,
		Price:          tourDTO.Price,
		Tags:           strings.Join(tourDTO.Tags, "|"),
		AuthorID:       tourDTO.AuthorID,
		Status:         tourDTO.Status,
		TourTimes:      "TODO",
		TourRatings:    "TODO",
		Closed:         tourDTO.Closed,
	}

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

	return tour, equipment, checkpoints
}