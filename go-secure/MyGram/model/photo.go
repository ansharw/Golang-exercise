package model

type Photo struct {
	GormModel
	Title    string `gorm:"not null" json:"title" form:"title" binding:"required" validate:"required"`
	Caption  string `json:"caption,omitempty" form:"caption,omitempty"`
	PhotoURL string `gorm:"not null" json:"photo_url" form:"photo_url" binding:"required" validate:"required,url"`
	UserID   uint   `gorm:"not null" json:"user_id" form:"user_id"`
}
