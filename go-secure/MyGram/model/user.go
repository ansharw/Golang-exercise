package model

type User struct {
	GormModel
	Username string `gorm:"not null,unique" json:"username" form:"username" binding:"required"`
	Email    string `gorm:"not null" json:"email" form:"email" binding:"required,email"`
	Password string `gorm:"not null" json:"password" form:"password" binding:"required,min=6"`
	Age      int    `gorm:"not null" json:"age" form:"age" binding:"required,gte=8,lte=130"`
}

type RequestUserLogin struct {
	Email    string `gorm:"-:all" json:"email" form:"email" binding:"required,email"`
	Password string `gorm:"-:all" json:"password" form:"password" binding:"required,min=6"`
}

// response
type ResponseErrorGeneral struct {
	Status  string `json:"error"`
	Message string `json:"message"`
}

type ResponseRegistered struct {
	Id  uint `json:"id"`
	Email string `json:"email"`
	Username string `json:"username"`
}

type ResponseDeleted struct {
	Message  string `json:"message"`
}

type ResponseToken struct {
	Token  string `json:"token"`
}
