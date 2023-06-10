package entity

type TipeMesin struct {
	ID    uint8 `json:"id" gorm:"primaryKey;autoIncrement"`
	Tipe  int   `json:"tipe" gorm:"not null"`
	Mobil []Mobil

	Timestamp
}

type TipeMobil struct {
	ID    uint8  `json:"id" gorm:"primaryKey;autoIncrement"`
	Tipe  string `json:"tipe" gorm:"not null"`
	Mobil []Mobil

	Timestamp
}

type TipePersneling struct {
	ID    uint8 `json:"id" gorm:"primaryKey;autoIncrement"`
	Tipe  int   `json:"tipe" gorm:"not null"`
	Mobil []Mobil

	Timestamp
}
