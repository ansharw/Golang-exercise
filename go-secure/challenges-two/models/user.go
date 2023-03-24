package models

import (
	"challenges-two/helpers"
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Fullname string    `gorm:"not null;autoIncrement;default:1" json:"full_name" form:"full_name" valid:"required~Your full name is required"`
	Email    string    `gorm:"not null" json:"email" form:"email" valid:"required~Invalid email format"`
	Password string    `gorm:"not null" json:"password" form:"password" valid:"required~Password has to have minimum length of 6 characters"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
