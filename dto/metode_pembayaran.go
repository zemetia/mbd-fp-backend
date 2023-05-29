package dto

import "github.com/google/uuid"

//Todo : Implement service, controller
type MetodePembayaranCreateDto struct {
	ID     uuid.UUID `gorm:"primary_key" json:"id" form:"id"`
	Metode string    `json:"metode" form:"metode" binding:"required"`
}

type MetodePembayaranUpdateDto struct {
	ID     uuid.UUID `gorm:"primary_key" json:"id" form:"id"`
	Metode string    `json:"metode" form:"metode"`
}
