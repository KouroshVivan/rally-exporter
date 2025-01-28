package models

type Workload struct {
	Model

	ID                   uint    `gorm:"primary_key;not null"`
	UUID                 string  `gorm:"type:varchar(36);not null"`
	TaskUUID             string  `gorm:"type:varchar(36);not null"`
	SubtaskUUID          string  `gorm:"type:varchar(36);not null"`
	Name                 string  `gorm:"type:text"`
	Description          string  `gorm:"type:text"`
	Position             float64 `gorm:"type:float"`
	Runner               string  `gorm:"type:text"`
	RunnerType           string  `gorm:"type:text"`
	Contexts             string  `gorm:"type:text"`
	ContextsResults      string  `gorm:"type:text"`
	Sla                  string  `gorm:"type:text"`
	SlaResults           string  `gorm:"type:text"`
	Args                 string  `gorm:"type:text"`
	Hooks                string  `gorm:"type:text"`
	StartTime            float64 `gorm:"type:float"`
	LoadDuration         float64 `gorm:"type:float"`
	FullDuration         float64 `gorm:"type:float"`
	MinDuration          float64 `gorm:"type:float"`
	MaxDuration          float64 `gorm:"type:float"`
	TotalIterationCount  float64 `gorm:"type:float"`
	FailedIterationCount float64 `gorm:"type:float"`
	Statistics           string  `gorm:"type:text"`
	PassSLA              bool
}
