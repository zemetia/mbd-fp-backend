package entity

import (
	"github.com/google/uuid"
)

type MetodePembayaran struct {
	ID     uuid.UUID `json:"id" gorm:"primary_key;not_null"`
	Metode string    `json:"metode" gorm:"size:20;not null"`

	Transaksi []Transaksi
}
