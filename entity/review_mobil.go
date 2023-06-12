package entity

import (
	"github.com/google/uuid"
)

type ReviewMobil struct {
	ID      uuid.UUID `json:"id" gorm:"primaryKey;size:36;not_null"`
	Review  string    `json:"review" gorm:"not null"`
	Rating  uint16    `json:"rating" gorm:"not null"`
	MobilID uuid.UUID `json:"mobil_id" gorm:"not null;size:36"`
	UserID  uuid.UUID `json:"user_id" gorm:"not null;size:36"`
	Timestamp
}
