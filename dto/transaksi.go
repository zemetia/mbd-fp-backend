package dto

import (
	"time"

	"github.com/google/uuid"
)

type TransaksiCreateDto struct {
	ID                 uuid.UUID `gorm:"primary_key;not_null"`
	TglAmbil           time.Time `json:"tgl_ambil" form:"tgl_ambil" binding:"required"`
	TglKembali         time.Time `json:"tgl_kembali" form:"tgl_kembali" binding:"required"`
	TglDikembalikan    time.Time `json:"tgl_dikembalikan" form:"tgl_dikembalikan" binding:"required"`
	Diskon             float32   `json:"diskon" form:"diskon" binding:"required"`
	TotalHarga         float32   `json:"total_harga" form:"total_harga" binding:"required"`
	TotalDenda         float32   `json:"total_denda" form:"total_denda" binding:"required"`
	MobilID            uuid.UUID `json:"mobil_id" form:"mobil_id" binding:"required"`
	UserID             uuid.UUID `json:"user_id" form:"user_id" binding:"required"`
	MetodePembayaranID uuid.UUID `json:"metode_pembayaran_id" form:"metode_pembayaran_id" binding:"required"`
	LokasiID           uuid.UUID `json:"lokasi_id" form:"lokasi_id" binding:"required"`
}

type TransaksiUpdateDto struct {
	ID                 uuid.UUID `gorm:"primary_key;not_null"`
	TglAmbil           time.Time `json:"tgl_ambil" form:"tgl_ambil" binding:"required"`
	TglKembali         time.Time `json:"tgl_kembali" form:"tgl_kembali" binding:"required"`
	TglDikembalikan    time.Time `json:"tgl_dikembalikan" form:"tgl_dikembalikan" binding:"required"`
	Diskon             float32   `json:"diskon" form:"diskon" binding:"required"`
	TotalHarga         float32   `json:"total_harga" form:"total_harga" binding:"required"`
	TotalDenda         float32   `json:"total_denda" form:"total_denda" binding:"required"`
	MobilID            uuid.UUID `json:"mobil_id" form:"mobil_id" binding:"required"`
	UserID             uuid.UUID `json:"user_id" form:"user_id" binding:"required"`
	MetodePembayaranID uuid.UUID `json:"metode_pembayaran_id" form:"metode_pembayaran_id" binding:"required"`
	LokasiID           uuid.UUID `json:"lokasi_id" form:"lokasi_id" binding:"required"`
}
