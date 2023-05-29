package entity

import (
	"github.com/google/uuid"
)

type Lokasi struct {
	ID         uuid.UUID `gorm:"primary_key;not_null" json:"id"`
	name       string    `gorm:"not_null" json:"name"`
	longtitude string    `gorm:"not_null" json:"longtitude"`
	latitude   string    `gorm:"not_null" json:"latitude"`
}
