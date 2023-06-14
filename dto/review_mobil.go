package dto

type ReviewCreateDto struct {
	ID      string `gorm:"primary_key" json:"id" form:"id"`
	Review  string `json:"review" form:"review" binding:"required"`
	Rating  uint16 `json:"rating" form:"rating" binding:"required"`
	MobilID string `json:"mobil_id" form:"mobil_id" binding:"required"`
	UserID  string `json:"user_id" form:"mobil_id" binding:"required"`
}

type ReviewUpdateDto struct {
	ID      string `gorm:"primary_key" json:"id" form:"id"`
	Review  string `json:"review" form:"review"`
	Rating  uint16 `json:"rating" form:"rating"`
	MobilID string `json:"mobil_id" form:"mobil_id" binding:"required"`
	UserID  string `json:"user_id" form:"mobil_id" binding:"required"`
}
