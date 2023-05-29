package dto

//Todo : Implement service, controller
type MembershipCreateDto struct {
	Tier   string `gorm:"primary_key" json:"tier" form:"tier"`
	Diskon uint16 `json:"diskon" form:"diskon" binding:"required"`
}

type MembershipUpdateDto struct {
	Tier   string `gorm:"primary_key" json:"tier" form:"tier"`
	Diskon uint16 `json:"diskon" form:"diskon"`
}
