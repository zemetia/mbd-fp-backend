package entity

import (
	"github.com/google/uuid"
)

type ReviewMobil struct {
	ID      uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Review  string    `json:"review" gorm:"not null"`
	Rating  uint16    `json:"rating" gorm:"not null"`
	MobilID uuid.UUID `json:"mobil_id" gorm:"not null"`

	Timestamp
}
