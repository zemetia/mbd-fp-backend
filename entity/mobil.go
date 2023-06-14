package entity

type Mobil struct {
	ID                 string  `json:"id" gorm:"primary_key;size:36;not_null"`
	Nama               string  `json:"nama" gorm:"size:20;not null"`
	Price              float32 `json:"price" gorm:"not null"`
	PelatNo            string  `json:"pelat_no" gorm:"not null"`
	KapasitasPenumpang uint8   `json:"kapasitas_penumpang" gorm:"not null"`
	Status             bool    `json:"status" gorm:"default:true"`
	KapasitasMesin     string  `json:"kapasitas_mesin" gorm:"not null"`
	PhotoURL           string  `json:"photo_url"`

	UserID           string  `json:"user_id" gorm:"size:36"`
	TipeMobilID      uint8   `json:"tipe_mobil_id"`
	TipePersnelingID uint8   `json:"tipe_persneling_id"`
	TipeMesinID      uint8   `json:"tipe_mesin_id"`
	LokasiID         string  `json:"lokasi_id" gorm:"size:36"`
	Rating           float32 `json:"rating" gorm:"not null"`

	ReviewMobil []ReviewMobil `json:",omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Transaksi   []Transaksi   `json:",omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
