package model

type Comment struct {
	GormModel
	Message string `gorm:"not null" json:"message" form:"message" binding:"required" validate:"required"`
	UserID  uint   `gorm:"not null" json:"user_id" form:"user_id"`
	PhotoID uint   `gorm:"not null" json:"photo_id" form:"photo_id"`
}
