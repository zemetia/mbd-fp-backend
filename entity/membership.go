package entity

type Membership struct {
	Tier   string `json:"tier" gorm:"size:16;primary_key;not_null"`
	Diskon int    `json:"diskon" gorm:"not null"`

	User []User `json:",omitempty" gorm:"foreignKey:MembershipTier;constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
}
