package dto

import (
	"github.com/google/uuid"
)

type LokasiCreateDto struct {
	ID         uuid.UUID `gorm:"primary_key" json:"id" form:"id"`
	Name       string    `json:"name" form:"name" binding:"required"`
	Longtitude string    `json:"longtitude" form:"longtitude" binding:"required"`
	Latitude   string    `json:"latitude" form:"latitude" binding:"required"`
}

type LokasiUpdateDto struct {
	ID         uuid.UUID `gorm:"primary_key" json:"id" form:"id"`
	Name       string    `json:"name" form:"name"`
	Longtitude string    `json:"longtitude" form:"longtitude"`
	Latitude   string    `json:"latitude" form:"latitude"`
}

type LokasiGetAllDto struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Longtitude string    `json:"longtitude"`
	Latitude   string    `json:"latitude"`
	Distance   float64   `json:"distance"`
}
