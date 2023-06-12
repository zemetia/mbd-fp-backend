package entity

import (
	"github.com/google/uuid"
)

type Lokasi struct {
	ID         uuid.UUID `gorm:"primary_key;size:36;not_null" json:"id"`
	Name       string    `gorm:"not_null" json:"name"`
	Longtitude string    `gorm:"not_null" json:"longtitude"`
	Latitude   string    `gorm:"not_null" json:"latitude"`

	Mobil     []Mobil
	Transaksi []Transaksi

	Timestamp
}
