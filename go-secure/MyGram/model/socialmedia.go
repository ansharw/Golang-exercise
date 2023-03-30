package model

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name" form:"name" binding:"required" validate:"required,unique"`
	SocialMediaURL string `gorm:"not null" json:"social_media_url" form:"social_media_url" binding:"required" validate:"required,url"`
	UserID         uint   `gorm:"not null" json:"user_id" form:"user_id"`
}
