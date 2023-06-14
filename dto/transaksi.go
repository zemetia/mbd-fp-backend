package dto

import (
	"time"

	"github.com/google/uuid"
)

type TransaksiCreateDto struct {
	ID         uuid.UUID `gorm:"primary_key;not_null"`
	TglAmbil   time.Time `json:"tgl_ambil" form:"tgl_ambil" binding:"required"`
	TglKembali time.Time `json:"tgl_kembali" form:"tgl_kembali" binding:"required"`
	// TglDikembalikan    time.Time `json:"tgl_dikembalikan" form:"tgl_dikembalikan"`
	// Diskon             float32   `json:"diskon" form:"diskon"`
	// TotalHarga         float32   `json:"total_harga" form:"total_harga"`
	// TotalDenda         float32   `json:"total_denda" form:"total_denda"`
	MobilID            uuid.UUID `json:"mobil_id" form:"mobil_id" binding:"required"`
	UserID             uuid.UUID `json:"user_id" form:"user_id" binding:"required"`
	MetodePembayaranID uint8     `json:"metode_pembayaran_id" form:"metode_pembayaran_id" binding:"required"`
	LokasiID           uuid.UUID `json:"lokasi_id" form:"lokasi_id" binding:"required"`
}

type TransaksiUpdateDto struct {
	ID                 uuid.UUID `gorm:"primary_key;not_null"`
	TglAmbil           time.Time `json:"tgl_ambil" form:"tgl_ambil" binding:"required"`
	TglKembali         time.Time `json:"tgl_kembali" form:"tgl_kembali" binding:"required"`
	TglDikembalikan    time.Time `json:"tgl_dikembalikan" form:"tgl_dikembalikan" binding:"required"`
	Diskon             float32   `json:"diskon" form:"diskon"`
	TotalHarga         float32   `json:"total_harga" form:"total_harga"`
	TotalDenda         float32   `json:"total_denda" form:"total_denda"`
	MobilID            uuid.UUID `json:"mobil_id" form:"mobil_id" binding:"required"`
	UserID             uuid.UUID `json:"user_id" form:"user_id" binding:"required"`
	MetodePembayaranID uint8     `json:"metode_pembayaran_id" form:"metode_pembayaran_id" binding:"required"`
	LokasiID           uuid.UUID `json:"lokasi_id" form:"lokasi_id" binding:"required"`
}

type TransaksiGetDto struct {
	ID                 uuid.UUID `gorm:"primary_key;not_null"`
	TglAmbil           time.Time `json:"tgl_ambil" form:"tgl_ambil" binding:"required"`
	TglKembali         time.Time `json:"tgl_kembali" form:"tgl_kembali" binding:"required"`
	TglDikembalikan    time.Time `json:"tgl_dikembalikan" form:"tgl_dikembalikan" binding:"required"`
	Diskon             float32   `json:"diskon" form:"diskon"`
	TotalHarga         float32   `json:"total_harga" form:"total_harga"`
	TotalDenda         float32   `json:"total_denda" form:"total_denda"`
	Mobil              string    `json:"mobil"`
	MobilID            uuid.UUID `json:"mobil_id" form:"mobil_id" binding:"required"`
	Mitra              string    `json:"mitra"`
	UserID             uuid.UUID `json:"user_id" form:"user_id" binding:"required"`
	MetodePembayaran   string    `json:"metode_pembayaran"`
	MetodePembayaranID uint8     `json:"metode_pembayaran_id" form:"metode_pembayaran_id" binding:"required"`
	Lokasi             string    `json:"lokasi"`
	LokasiID           uuid.UUID `json:"lokasi_id" form:"lokasi_id" binding:"required"`
}
