package entity

type MetodePembayaran struct {
	ID     uint8  `json:"id" gorm:"primary_key;not_null;autoIncrement"`
	Metode string `json:"metode" gorm:"size:20;not null"`

	Transaksi []Transaksi
}
