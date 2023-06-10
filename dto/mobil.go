package dto

import (
	"github.com/google/uuid"
)

type MobilCreateDto struct {
	ID             uuid.UUID `gorm:"primary_key;not_null"`
	Nama           string    `json:"nama" form:"nama" binding:"required"`
	Price          float32   `json:"price" form:"price" binding:"required"`
	PelatNo        string    `json:"pelat" form:"pelat" binding:"required"`
	Status         bool      `json:"status" form:"status" binding:"required"`
	KapasitasMesin string    `json:"kapasitas_mesin" form:"kapasitas_mesin" binding:"required"`

	UserID           uuid.UUID `json:"mitra_id" form:"mitra_id" binding:"required"`
	TipeMobilID      uint      `json:"tipe_mobil_id" form:"tipe_mobil_id" binding:"required"`
	TipePersnelingID uint      `json:"tipe_persneling_id" form:"tipe_persneling_id" binding:"required"`
	TipeMesinID      uint      `json:"tipe_mesin_id" form:"tipe_mesin_id" binding:"required"`
	LokasiID         uuid.UUID `json:"lokasi_id" form:"lokasi_id" binding:"required"`
}

type MobilUpdateDto struct {
	ID             uuid.UUID `gorm:"primary_key;not_null"`
	Nama           string    `json:"nama" form:"nama" binding:"required"`
	Price          float32   `json:"price" form:"price" binding:"required"`
	PelatNo        string    `json:"pelat" form:"pelat" binding:"required"`
	Status         bool      `json:"status" form:"status" binding:"required"`
	KapasitasMesin string    `json:"kapasitas_mesin" form:"kapasitas_mesin" binding:"required"`

	UserID           uuid.UUID `json:"mitra_id" form:"mitra_id" binding:"required"`
	TipeMobilID      uint      `json:"tipe_mobil_id" form:"tipe_mobil_id" binding:"required"`
	TipePersnelingID uint      `json:"tipe_persneling_id" form:"tipe_persneling_id" binding:"required"`
	TipeMesinID      uint      `json:"tipe_mesin_id" form:"tipe_mesin_id" binding:"required"`
	LokasiID         uuid.UUID `json:"lokasi_id" form:"lokasi_id" binding:"required"`
}
