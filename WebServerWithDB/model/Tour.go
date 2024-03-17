package model

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

