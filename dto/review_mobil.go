package dto

import "github.com/google/uuid"

//Todo : Implement service, controller
type ReviewCreateDto struct {
	ID     uuid.UUID `gorm:"primary_key" json:"id" form:"id"`
	Review string    `json:"review" form:"review" binding:"required"`
	Rating uint16    `json:"rating" form:"rating" binding:"required"`
}

type ReviewUpdateDto struct {
	ID     uuid.UUID `gorm:"primary_key" json:"id" form:"id"`
	Review string    `json:"review" form:"review"`
	Rating uint16    `json:"rating" form:"rating"`
}
