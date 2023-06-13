package dto

//Todo : Implement service, controller
type MetodePembayaranCreateDto struct {
	ID     uint8  `gorm:"primary_key" json:"id" form:"id"`
	Metode string `json:"metode" form:"metode" binding:"required"`
}

type MetodePembayaranUpdateDto struct {
	ID     uint8  `gorm:"primary_key" json:"id" form:"id"`
	Metode string `json:"metode" form:"metode"`
}
