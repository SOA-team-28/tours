package model

import "time"

type ExecutionStatus int

const (
	Completed ExecutionStatus = iota
	Abandoned
	InProgress
)

type Equipment struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Checkpoint struct {
	ID                   int     `json:"id"`
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

type Tour struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	DemandingLevel string `json:"demandignessLevel"`
	Price          int    `json:"price"`
	Tags           string `json:"tags"`
	AuthorID       int    `json:"authorId"`
	Status         string `json:"status"`
	Equipment      string `json:"equipment"`
	Checkpoints    string `json:"checkpoints"`
	TourTimes      string `json:"tourTimes"`
	TourRatings    string `json:"tourRatings"`
	Closed         bool   `json:"closed"`
}

type TourRatingPreview struct {
	ID           int       `json:"id"`
	Rating       int       `json:"rating"`
	Comment      string    `json:"comment"`
	TouristID    int       `json:"touristId"`
	TourID       int       `json:"tourId"`
	TourDate     time.Time `json:"tourDate"`
	CreationDate time.Time `json:"creationDate"`
	ImageNames   []string  `json:"imageNames"`
}

type CheckpointCompletition struct {
	ID               int       `json:"id"`
	TourExecutionID  int       `json:"touristId"`
	CheckpointID     int       `json:"tourId"`
	CompletitionTime time.Time `json:"start"`
}

type TourExecution struct {
	ID                   int             `json:"id"`
	TouristID            int             `json:"touristId"`
	TourID               int             `json:"tourId"`
	Start                time.Time       `json:"start"`
	LastActivity         time.Time       `json:"lastActivity"`
	ExecutionStatus      ExecutionStatus `json:"executionStatus"`
	CompletedCheckpoints []int           `json:"completedCheckpoints"`
}
