package entity

import (
	"github.com/google/uuid"
)

type MetodePembayaran struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
	Metode string    `gorm:"size:20;not null"`
}
