package model

type Comments struct {
	GormModel
	Message string `gorm:"not null" json:"message" form:"message" binding:"required"`
	UserID  uint   `gorm:"not null" json:"user_id" form:"user_id"`
	PhotoID uint   `gorm:"not null" json:"photo_id" form:"photo_id" binding:"required"`
}

// create, update
type RequestComments struct {
	Message string `gorm:"-:all" json:"message" form:"message" binding:"required"`
	PhotoID uint   `gorm:"-:all" json:"photo_id" form:"photo_id" binding:"required"`
}

type RequestGetComments struct {
	PhotoID uint   `gorm:"-:all" json:"photo_id" form:"photo_id" binding:"required"`
}

type RequestDeleteComments struct {
	PhotoID uint   `gorm:"-:all" json:"photo_id" form:"photo_id" binding:"required"`
}