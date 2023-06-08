package dto

import (
	"fp-mbd-amidrive/entity"

	"github.com/google/uuid"
)

type MobilCreateDto struct {
	ID             uuid.UUID `gorm:"primary_key;not_null"`
	Nama           string    `json:"nama" binding:"required"`
	Price          float32   `json:"price" binding:"required"`
	PelatNo        string    `json:"pelat" binding:"required"`
	Status         bool      `json:"status" binding:"required"`
	KapasitasMesin string    `json:"kapasitas_mesin" binding:"required"`

	MitraID          uuid.UUID `json:"mitra_id" binding:"required" gorm:"not null"`
	TipeMobilID      uint      `json:"tipe_mobil_id" binding:"required"`
	TipePersnelingID uint      `json:"tipe_persneling_id" binding:"required"`
	TipeMesinID      uint      `json:"tipe_mesin_id" binding:"required"`

	Lokasi entity.Lokasi
}

type MobilUpdateDto struct {
	ID             uuid.UUID `gorm:"primary_key;not_null"`
	Nama           string    `json:"nama"`
	Price          float32   `json:"price"`
	PelatNo        string    `json:"pelat"`
	Status         bool      `json:"status"`
	KapasitasMesin string    `json:"kapasitas_mesin"`

	MitraID          uuid.UUID `json:"mitra_id" gorm:"not null"`
	TipeMobilID      uint      `json:"tipe_mobil_id"`
	TipePersnelingID uint      `json:"tipe_persneling_id"`
	TipeMesinID      uint      `json:"tipe_mesin_id"`

	Lokasi entity.Lokasi
}
