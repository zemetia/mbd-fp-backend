package entity

type TipeMesin struct {
	ID   uint8 `gorm:"primaryKey;autoIncrement"`
	Tipe int   `gorm:"not null"`

	Timestamp
}

type TipeMobil struct {
	ID   uint8  `gorm:"primaryKey;autoIncrement"`
	Tipe string `gorm:"not null"`

	Timestamp
}

type TipePersneling struct {
	ID   uint8 `gorm:"primaryKey;autoIncrement"`
	Tipe int   `gorm:"not null"`

	Timestamp
}
