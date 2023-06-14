package entity

import (
	"time"

	"github.com/google/uuid"
)

type Transaksi struct {
	ID                 uuid.UUID `json:"id" gorm:"primary_key;size:36;not_null"`
	TglAmbil           time.Time `json:"tgl_ambil" gorm:"not null"`
	TglKembali         time.Time `json:"tgl_kembali" gorm:"not null"`
	TglDikembalikan    time.Time `json:"tgl_dikembalikan" gorm:"default:null"`
	Diskon             float32   `json:"diskon"`
	TotalHarga         float32   `json:"total_harga"`
	TotalDenda         float32   `json:"total_denda"`
	MobilID            uuid.UUID `json:"mobil_id" gorm:"size:36;"`
	User               User
	UserID             uuid.UUID `json:"user_id" gorm:"size:36;"`
	MetodePembayaranID uint8     `json:"metode_pembayaran_id"`
	LokasiID           uuid.UUID `json:"lokasi_id" gorm:"size:36"`

	Timestamp
}
