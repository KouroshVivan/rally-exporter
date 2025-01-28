package models

import "time"

type Workloaddata struct {
	Model

	ID                   uint    `gorm:"primary_key;not null"`
	UUID                 string  `gorm:"type:varchar(36);not null"`
	TaskUUID             string  `gorm:"type:varchar(36);not null"`
	WorkloadUUID         string  `gorm:"type:varchar(36);not null"`
	ChunkOrder           float64 `gorm:"type:float"`
	ChunkData            string  `gorm:"type:text"`
	IterationCount       float64 `gorm:"type:float"`
	FailedIterationCount float64 `gorm:"type:float"`
	ChunkSize            float64 `gorm:"type:float"`
	CompressedChunkSize  float64 `gorm:"type:float"`
	StartedAt            time.Time
	FinishedAt           time.Time
}
