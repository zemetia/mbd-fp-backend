package entity

import (
	"time"

	"github.com/google/uuid"
)

type Transaksi struct {
	ID                 uuid.UUID `json:"id" gorm:"primary_key;not_null"`
	TglAmbil           time.Time `json:"tgl_ambil" gorm:"not null"`
	TglKembali         time.Time `json:"tgl_kembali" gorm:"not null"`
	TglDikembalikan    time.Time `json:"tgl_dikembalikan" gorm:"not null"`
	Diskon             float32   `json:"diskon" gorm:"not null"`
	TotalHarga         float32   `json:"total_harga" gorm:"not null"`
	TotalDenda         float32   `json:"total_denda" gorm:"not null"`
	MobilID            uuid.UUID `json:"mobil_id" gorm:"not null"`
	UserID             uuid.UUID `json:"user_id" gorm:"not null"`
	MetodePembayaranID uuid.UUID `json:"metode_pembayaran_id" gorm:"size:36;not null"`
	LokasiID           uuid.UUID `json:"lokasi_id" gorm:"size:36"`

	Timestamp
}
