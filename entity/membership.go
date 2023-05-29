package entity

type Membership struct {
	Tier   string `gorm:"size:16;primaryKey;not null"`
	Diskon int    `gorm:"not null"`
}
