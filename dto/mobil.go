package dto

type MobilCreateDto struct {
	// ID                 string  `gorm:"primary_key;not_null"`
	Nama               string  `json:"nama" form:"nama" binding:"required"`
	Price              float32 `json:"price" form:"price" binding:"required"`
	PelatNo            string  `json:"pelat_no" form:"pelat_no" binding:"required"`
	KapasitasPenumpang uint8   `json:"kapasitas_penumpang" form:"kapasitas_penumpang" binding:"required"`
	Status             bool    `json:"status" form:"status" binding:"required"`
	KapasitasMesin     string  `json:"kapasitas_mesin" form:"kapasitas_mesin" binding:"required"`
	PhotoURL           string  `json:"photo_url" form:"photo_url"`

	UserID           string  `json:"user_id" form:"user_id" binding:"required"`
	TipeMobilID      uint8   `json:"tipe_mobil_id" form:"tipe_mobil_id" binding:"required"`
	TipePersnelingID uint8   `json:"tipe_persneling_id" form:"tipe_persneling_id" binding:"required"`
	TipeMesinID      uint8   `json:"tipe_mesin_id" form:"tipe_mesin_id" binding:"required"`
	LokasiID         string  `json:"lokasi_id" form:"lokasi_id" binding:"required"`
	Rating           float32 `json:"rating" form:"rating" binding:"required"`
}

type MobilUpdateDto struct {
	ID                 string  `gorm:"primary_key;not_null"`
	Nama               string  `json:"nama" form:"nama" binding:"required"`
	Price              float32 `json:"price" form:"price" binding:"required"`
	PelatNo            string  `json:"pelat" form:"pelat" binding:"required"`
	KapasitasPenumpang uint8   `json:"kapasitas_penumpang" form:"kapasitas_penumpang" binding:"required"`
	Status             bool    `json:"status" form:"status" binding:"required"`
	KapasitasMesin     string  `json:"kapasitas_mesin" form:"kapasitas_mesin" binding:"required"`
	PhotoURL           string  `json:"photo_url" form:"photo_url"`

	UserID           string  `json:"mitra_id" form:"mitra_id" binding:"required,uuid"`
	TipeMobilID      uint8   `json:"tipe_mobil_id" form:"tipe_mobil_id" binding:"required"`
	TipePersnelingID uint8   `json:"tipe_persneling_id" form:"tipe_persneling_id" binding:"required"`
	TipeMesinID      uint8   `json:"tipe_mesin_id" form:"tipe_mesin_id" binding:"required"`
	LokasiID         string  `json:"lokasi_id" form:"lokasi_id" binding:"required,uuid"`
	Rating           float32 `json:"rating" form:"rating" binding:"required"`
}

type MobilGetDto struct {
	ID                 string
	Nama               string  `json:"nama"`
	Price              float32 `json:"price"`
	PelatNo            string  `json:"pelat"`
	KapasitasPenumpang uint8   `json:"kapasitas_penumpang"`
	Status             bool    `json:"status"`
	KapasitasMesin     string  `json:"kapasitas_mesin"`
	PhotoURL           string  `json:"photo_url"`

	Mitra            string  `json:"mitra"`
	MitraID          string  `json:"mitra_id,omitempty"`
	TipeMobil        string  `json:"tipe_mobil"`
	TipeMobilID      uint    `json:"tipe_mobil_id,omitempty"`
	TipePersneling   string  `json:"tipe_persneling"`
	TipePersnelingID uint    `json:"tipe_persneling_id,omitempty"`
	TipeMesin        string  `json:"tipe_mesin"`
	TipeMesinID      uint    `json:"tipe_mesin_id,omitempty"`
	Lokasi           string  `json:"lokasi"`
	LokasiID         string  `json:"lokasi_id,omitempty"`
	Distance         float64 `json:"distance,omitempty"`
	Rating           float32 `json:"rating"`
}
