package entity

import (
	"github.com/google/uuid"
)

type Mobil struct {
	ID             uuid.UUID `json:"id" gorm:"primary_key;not_null"`
	Nama           string    `json:"nama" gorm:"size:20;not null"`
	Price          float32   `json:"price" gorm:"not null"`
	PelatNo        string    `json:"pelat_no" gorm:"size:10;not null"`
	Status         bool      `json:"status" gorm:"default:true"`
	KapasitasMesin string    `json:"kapasitas_mesin" gorm:"not null"`

	MitraID          uuid.UUID `json:"mitra_id" gorm:"not null"`
	TipeMobilID      uint      `json:"tipe_mobil_id" gorm:"not null"`
	TipePersnelingID uint      `json:"tipe_persneling_id" gorm:"not null"`
	TipeMesinID      uint      `json:"tipe_mesin_id" gorm:"not null"`

	LokasiID uuid.UUID `json:"lokasi_id" gorm:"size:36"`
	Lokasi   Lokasi    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
