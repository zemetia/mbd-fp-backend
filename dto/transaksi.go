package dto

import (
	"time"

	"github.com/google/uuid"
)

type TransaksiCreateDto struct {
	ID                 uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
	TglAmbil           time.Time `json:"tgl_ambil" form:"tgl_ambil" binding:"required"`
	TglKembali         time.Time `json:"tgl_kembali" form:"tgl_kembali" binding:"required"`
	Diskon             float32   `json:"diskon" form:"diskon" binding:"required"`
	MobilID            uuid.UUID `json:"mobil_id" form:"mobil_id" binding:"required"`
	PelangganID        uuid.UUID `json:"pelanggan_id" form:"pelanggan_id" binding:"required"`
	LokasiID           uuid.UUID `json:"lokasi_id" form:"lokasi_id" binding:"required"`
	MetodePembayaranID uuid.UUID `json:"metode_pembayaran_id" form:"metode_pembayaran_id" binding:"required"`
}

type TransaksiUpdateDto struct {
	ID                 uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
	TglAmbil           time.Time `json:"tgl_ambil" form:"tgl_ambil" binding:"required"`
	TglKembali         time.Time `json:"tgl_kembali" form:"tgl_kembali" binding:"required"`
	Diskon             float32   `json:"diskon" form:"diskon" binding:"required"`
	MobilID            uuid.UUID `json:"mobil_id" form:"mobil_id" binding:"required"`
	PelangganID        uuid.UUID `json:"pelanggan_id" form:"pelanggan_id" binding:"required"`
	LokasiID           uuid.UUID `json:"lokasi_id" form:"lokasi_id" binding:"required"`
	MetodePembayaranID uuid.UUID `json:"metode_pembayaran_id" form:"metode_pembayaran_id" binding:"required"`
}
