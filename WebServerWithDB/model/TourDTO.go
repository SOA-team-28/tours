package model

import "strings"

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
	Checkpoints    []Checkpoint `json:"checkpoints"`
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

	return tour, tourDTO.Equipment, tourDTO.Checkpoints
}