package dto

import "github.com/google/uuid"

type ReviewCreateDto struct {
	ID      uuid.UUID `gorm:"primary_key" json:"id" form:"id"`
	Review  string    `json:"review" form:"review" binding:"required"`
	Rating  uint16    `json:"rating" form:"rating" binding:"required"`
	MobilID uuid.UUID `json:"mobil_id" form:"mobil_id" binding:"required"`
	UserID  uuid.UUID `json:"user_id" form:"mobil_id" binding:"required"`
}

type ReviewUpdateDto struct {
	ID      uuid.UUID `gorm:"primary_key" json:"id" form:"id"`
	Review  string    `json:"review" form:"review"`
	Rating  uint16    `json:"rating" form:"rating"`
	MobilID uuid.UUID `json:"mobil_id" form:"mobil_id" binding:"required"`
	UserID  uuid.UUID `json:"user_id" form:"mobil_id" binding:"required"`
}
