package entity

import (
	"time"
)

type Transaksi struct {
	ID                 string            `json:"id" gorm:"primary_key;size:36;not_null"`
	TglAmbil           time.Time         `json:"tgl_ambil" gorm:"not null"`
	TglKembali         time.Time         `json:"tgl_kembali" gorm:"not null"`
	TglDikembalikan    time.Time         `json:"tgl_dikembalikan" gorm:"default:null"`
	Diskon             float32           `json:"diskon"`
	TotalHarga         float32           `json:"total_harga"`
	TotalDenda         float32           `json:"total_denda"`
	MobilID            string            `json:"mobil_id" gorm:"size:36;"`
	Mobil              *Mobil            `json:",omitempty"`
	UserID             string            `json:"user_id" gorm:"size:36;"`
	User               *User             `json:",omitempty"`
	MetodePembayaranID uint8             `json:"metode_pembayaran_id"`
	MetodePembayaran   *MetodePembayaran `json:",omitempty"`
	LokasiID           string            `json:"lokasi_id" gorm:"size:36"`

	Timestamp
}
