package model

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name" form:"name" binding:"required"`
	SocialMediaURL string `gorm:"not null" json:"social_media_url" form:"social_media_url" binding:"required,url"`
	UserID         uint   `gorm:"not null" json:"user_id" form:"user_id"`
}

// create, update
type RequestSocialMedia struct {
	Name           string `gorm:"-:all" json:"name" form:"name" binding:"required"`
	SocialMediaURL string `gorm:"-:all" json:"social_media_url" form:"social_media_url" binding:"required,url"`
}
