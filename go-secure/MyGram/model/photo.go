package model

type Photo struct {
	GormModel
	Title    string `gorm:"not null" json:"title" form:"title" binding:"required"`
	Caption  string `gorm:"not null" json:"caption,omitempty" form:"caption,omitempty"`
	PhotoURL string `gorm:"not null" json:"photo_url" form:"photo_url" binding:"required,url"`
	UserID   uint   `gorm:"not null" json:"user_id" form:"user_id"`
}

// create, update
type RequestPhoto struct {
	Title    string `gorm:"-:all" json:"title" form:"title" binding:"required"`
	Caption  string `gorm:"-:all" json:"caption,omitempty" form:"caption,omitempty"`
	PhotoURL string `gorm:"-:all" json:"photo_url" form:"photo_url" binding:"required,url"`
}