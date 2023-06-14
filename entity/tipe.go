package entity

type TipeMesin struct {
	ID    uint8   `json:"id" gorm:"primaryKey;autoIncrement"`
	Tipe  string  `json:"tipe" gorm:"not null"`
	Mobil []Mobil `json:",omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Timestamp
}

type TipeMobil struct {
	ID    uint8   `json:"id" gorm:"primaryKey;autoIncrement"`
	Tipe  string  `json:"tipe" gorm:"not null"`
	Mobil []Mobil `json:",omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Timestamp
}

type TipePersneling struct {
	ID    uint8   `json:"id" gorm:"primaryKey;autoIncrement"`
	Tipe  string  `json:"tipe" gorm:"not null"`
	Mobil []Mobil `json:",omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Timestamp
}
