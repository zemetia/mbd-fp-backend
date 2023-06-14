package dto

import (
	"time"
)

type TransaksiCreateDto struct {
	ID         string    `gorm:"primary_key;not_null"`
	TglAmbil   time.Time `json:"tgl_ambil" form:"tgl_ambil" binding:"required"`
	TglKembali time.Time `json:"tgl_kembali" form:"tgl_kembali" binding:"required"`
	// TglDikembalikan    time.Time `json:"tgl_dikembalikan" form:"tgl_dikembalikan"`
	// Diskon             float32   `json:"diskon" form:"diskon"`
	// TotalHarga         float32   `json:"total_harga" form:"total_harga"`
	// TotalDenda         float32   `json:"total_denda" form:"total_denda"`
	MobilID            string `json:"mobil_id" form:"mobil_id" binding:"required"`
	UserID             string `json:"user_id" form:"user_id" binding:"required"`
	MetodePembayaranID uint8  `json:"metode_pembayaran_id" form:"metode_pembayaran_id" binding:"required"`
	LokasiID           string `json:"lokasi_id" form:"lokasi_id" binding:"required"`
}

type TransaksiUpdateDto struct {
	ID                 string    `gorm:"primary_key;not_null"`
	TglAmbil           time.Time `json:"tgl_ambil" form:"tgl_ambil" binding:"required"`
	TglKembali         time.Time `json:"tgl_kembali" form:"tgl_kembali" binding:"required"`
	TglDikembalikan    time.Time `json:"tgl_dikembalikan" form:"tgl_dikembalikan" binding:"required"`
	Diskon             float32   `json:"diskon" form:"diskon"`
	TotalHarga         float32   `json:"total_harga" form:"total_harga"`
	TotalDenda         float32   `json:"total_denda" form:"total_denda"`
	MobilID            string    `json:"mobil_id" form:"mobil_id" binding:"required"`
	UserID             string    `json:"user_id" form:"user_id" binding:"required"`
	MetodePembayaranID uint8     `json:"metode_pembayaran_id" form:"metode_pembayaran_id" binding:"required"`
	LokasiID           string    `json:"lokasi_id" form:"lokasi_id" binding:"required"`
}

type TransaksiGetDto struct {
	ID                 string    `gorm:"primary_key;not_null"`
	TglAmbil           time.Time `json:"tgl_ambil" form:"tgl_ambil" binding:"required"`
	TglKembali         time.Time `json:"tgl_kembali" form:"tgl_kembali" binding:"required"`
	TglDikembalikan    time.Time `json:"tgl_dikembalikan" form:"tgl_dikembalikan" binding:"required"`
	Diskon             float32   `json:"diskon" form:"diskon"`
	TotalHarga         float32   `json:"total_harga" form:"total_harga"`
	TotalDenda         float32   `json:"total_denda" form:"total_denda"`
	Mobil              string    `json:"mobil"`
	MobilID            string    `json:"mobil_id" form:"mobil_id" binding:"required"`
	Mitra              string    `json:"mitra"`
	UserID             string    `json:"user_id" form:"user_id" binding:"required"`
	MetodePembayaran   string    `json:"metode_pembayaran"`
	MetodePembayaranID uint8     `json:"metode_pembayaran_id" form:"metode_pembayaran_id" binding:"required"`
	Lokasi             string    `json:"lokasi"`
	LokasiID           string    `json:"lokasi_id" form:"lokasi_id" binding:"required"`
}
