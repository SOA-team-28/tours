package model

import (
	"strings"
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
	RequiredTime         float64  `json:"requiredTimeInSeconds"`
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
	Price          float64         `json:"price"`
	Tags           []string        `json:"tags"`
	AuthorID       int             `json:"authorId"`
	Status         string          `json:"status"`
	Equipment      []Equipment     `json:"equipment"`
	Checkpoints    []CheckpointDTO `json:"checkpoints"`
	TourTimes      []int           `json:"tourTimes"`
	TourRatings    []int           `json:"tourRatings"`
	Closed         bool            `json:"closed"`
}

type MapObjectDTO struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	PictureURL  string  `json:"pictureURL"`
	Category    string  `json:"category"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
}

type TourPreferenceDTO struct {
	ID         int      `json:"id"`
	CreatorId  int      `json:"creatorId"`
	Difficulty string   `json:"difficulty"`
	Walk       int      `json:"walk"`
	Bike       int      `json:"bike"`
	Car        int      `json:"car"`
	Boat       int      `json:"boat"`
	Tags       []string `json:"tags"`
}

func (tourPreferenceDto *TourPreferenceDTO) MapToTourPreference() TourPreference {
	var tourDifficulty TourDifficulty

	switch tourPreferenceDto.Difficulty {
	case "Easy":
		tourDifficulty = Easy
	case "Medium":
		tourDifficulty = Medium
	default:
		tourDifficulty = Hard
	}

	tourPreference := TourPreference{
		ID:         tourPreferenceDto.ID,
		CreatorId:  tourPreferenceDto.CreatorId,
		Difficulty: tourDifficulty,
		Walk:       tourPreferenceDto.Walk,
		Bike:       tourPreferenceDto.Bike,
		Car:        tourPreferenceDto.Car,
		Boat:       tourPreferenceDto.Bike,
		Tags:       strings.Join(tourPreferenceDto.Tags, "|"),
	}

	return tourPreference
}

func (mapObjectDto *MapObjectDTO) MapToMapObject() MapObject {
	var category MapObjectType

	switch mapObjectDto.Category {
	case "Other":
		category = Other
	case "Restaurant":
		category = Restaurant
	case "WC":
		category = WC
	default:
		category = Parking
	}

	mapObject := MapObject{
		ID:          mapObjectDto.ID,
		Name:        mapObjectDto.Name,
		Description: mapObjectDto.Description,
		PictureURL:  mapObjectDto.PictureURL,
		Category:    category,
		Longitude:   mapObjectDto.Longitude,
		Latitude:    mapObjectDto.Latitude,
	}

	return mapObject
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
