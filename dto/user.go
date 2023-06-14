package dto

// Todo : add membership id (tier)
type UserCreateDto struct {
	ID             string `gorm:"primary_key" json:"id" form:"id"`
	Name           string `json:"name" form:"name" binding:"required"`
	Email          string `json:"email" form:"email" binding:"required"`
	NoTelp         string `json:"no_telp" form:"no_telp" binding:"required"`
	Password       string `json:"password" form:"password" binding:"required"`
	MembershipTier string `json:"membership" form:"membership" binding:"required"`
	PhotoURL       string `json:"photo_url" form:"photo_url"`
}

type UserUpdateDto struct {
	ID             string `gorm:"primary_key" json:"id" form:"id"`
	Name           string `json:"name" form:"name"`
	Email          string `json:"email" form:"email"`
	NoTelp         string `json:"no_telp" form:"no_telp"`
	Password       string `json:"password" form:"password"`
	MembershipTier string `json:"membership" form:"membership" binding:"required"`
	PhotoURL       string `json:"photo_url" form:"photo_url"`
}

type UserLoginDto struct {
	Email    string `json:"email" form:"email" binding:"email"`
	Password string `json:"password" form:"password" binding:"required"`
}
