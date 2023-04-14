package model

type User struct {
	GormModel
	Username string `gorm:"not null;unique" json:"username" form:"username" binding:"required"`
	Email    string `gorm:"not null;unique" json:"email" form:"email" binding:"required,email"`
	Password string `gorm:"not null" json:"password" form:"password" binding:"required,min=6"`
	Age      int    `gorm:"not null" json:"age" form:"age" binding:"required,gte=8"`
}

type RequestUserLogin struct {
	Email    string `gorm:"-:all" json:"email" form:"email" binding:"required,email"`
	Password string `gorm:"-:all" json:"password" form:"password" binding:"required"`
}

type RequestUserRegister struct {
	Age      int    `gorm:"-:all" json:"age" form:"age" binding:"required,gte=8"`
	Username string `gorm:"-:all;unique" json:"username" form:"username" binding:"required"`
	Email    string `gorm:"-:all;unique" json:"email" form:"email" binding:"required,email"`
	Password string `gorm:"-:all" json:"password" form:"password" binding:"required,min=6"`
}

// response
type ResponseErrorGeneral struct {
	Status  string `json:"error"`
	Message string `json:"message"`
}

type ResponseRegistered struct {
	Message string `json:"message"`
}

type ResponseDeleted struct {
	Message string `json:"message"`
}

type ResponseToken struct {
	Token string `json:"token"`
}
