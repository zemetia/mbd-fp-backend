package dto

type CreateTipeMesin struct {
	ID   uint8 `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
	Tipe int   `json:"tipe" form:"tipe" binding:"tipe"`
}

type UpdateTipeMesin struct {
	ID   uint8 `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
	Tipe int   `json:"tipe" form:"tipe"`
}

type CreateTipeMobil struct {
	ID   uint8  `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
	Tipe string `json:"tipe" form:"tipe" binding:"tipe"`
}

type UpdateTipeMobil struct {
	ID   uint8  `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
	Tipe string `json:"tipe" form:"tipe"`
}

type CreateTipePersneling struct {
	ID   uint8 `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
	Tipe int   `json:"tipe" form:"tipe" binding:"tipe"`
}

type UpdateTipePersneling struct {
	ID   uint8 `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
	Tipe int   `json:"tipe" form:"tipe"`
}
