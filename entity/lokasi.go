package entity

type Lokasi struct {
	ID         string `gorm:"primary_key;size:36;not_null" json:"id"`
	Name       string `gorm:"not_null" json:"name"`
	Longtitude string `gorm:"not_null" json:"longtitude"`
	Latitude   string `gorm:"not_null" json:"latitude"`

	Mobil     []Mobil     `json:",omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Transaksi []Transaksi `json:",omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Timestamp
}
