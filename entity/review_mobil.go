package entity

type ReviewMobil struct {
	ID      string `json:"id" gorm:"primaryKey;size:36;not_null"`
	Review  string `json:"review" gorm:"not null"`
	Rating  uint16 `json:"rating" gorm:"not null"`
	MobilID string `json:"mobil_id" gorm:"not null;size:36"`
	UserID  string `json:"user_id" gorm:"size:36"`
	Timestamp
}
