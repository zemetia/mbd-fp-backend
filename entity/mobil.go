package entity

import (
	"github.com/google/uuid"
)

type Mobil struct {
	ID               uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
	Nama             string    `gorm:"size:20;not null"`
	Lokasi           string    `gorm:"size:20;not null"`
	Price            int       `gorm:"not null"`
	PelatNo          string    `gorm:"size:10;not null"`
	Status           bool      `gorm:"not null"`
	KapasitasMesin   int       `gorm:"not null"`
	MitraID          uuid.UUID `gorm:"type:uuid;not null"`
	TipeMobilID      uint      `gorm:"not null"`
	TipePersnelingID uint      `gorm:"not null"`
	TipeMesinID      uint      `gorm:"not null"`
}
