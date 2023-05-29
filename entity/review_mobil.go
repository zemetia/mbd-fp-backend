package entity

import (
	"github.com/google/uuid"
)

type ReviewMobil struct {
	ID      uint      `gorm:"primaryKey;autoIncrement"`
	Review  string    `gorm:"not null"`
	MobilID uuid.UUID `gorm:"type:uuid;not null"`

	Timestamp
}
