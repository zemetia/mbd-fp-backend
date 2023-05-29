package entity

import (
	"time"

	"github.com/google/uuid"
)

type Transaksi struct {
	ID                 uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
	TglAmbil           time.Time `gorm:"not null"`
	TglKembali         time.Time `gorm:"not null"`
	TglDikembalikan    time.Time `gorm:"not null"`
	Diskon             float32   `gorm:"not null"`
	TotalHarga         float32   `gorm:"not null"`
	MobilID            uuid.UUID `gorm:"type:uuid;not null"`
	PelangganID        uuid.UUID `gorm:"type:uuid;not null"`
	LokasiID           uuid.UUID `gorm:"type:uuid;not null"`
	MetodePembayaranID uuid.UUID `gorm:"type:uuid;not null"`
	DendaID            uuid.UUID `gorm:"type:uuid"`

	Timestamp
}
