package entity

import (
	"github.com/google/uuid"
)

type Mobil struct {
	ID                 uuid.UUID `json:"id" gorm:"primary_key;size:36;not_null"`
	Nama               string    `json:"nama" gorm:"size:20;not null"`
	Price              float32   `json:"price" gorm:"not null"`
	PelatNo            string    `json:"pelat_no" gorm:"size:10;not null"`
	KapasitasPenumpang uint8     `json:"kapasitas_penumpang" gorm:"not null"`
	Status             bool      `json:"status" gorm:"default:true"`
	KapasitasMesin     string    `json:"kapasitas_mesin" gorm:"not null"`
	PhotoURL           string    `json:"photo_url"`

	UserID           uuid.UUID `json:"user_id" gorm:"size:36"`
	TipeMobilID      uint8     `json:"tipe_mobil_id"`
	TipePersnelingID uint8     `json:"tipe_persneling_id"`
	TipeMesinID      uint8     `json:"tipe_mesin_id"`
	LokasiID         uuid.UUID `json:"lokasi_id" gorm:"size:36;not null"`
	Rating           float32   `json:"rating" gorm:"not null"`

	ReviewMobil []ReviewMobil
}
