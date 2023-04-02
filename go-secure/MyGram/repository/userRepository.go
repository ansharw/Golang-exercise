package repository

import (
	"MyGram/model"
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	Login(ctx context.Context, tx *gorm.DB, email, pass string) (model.User, error)
	Register(ctx context.Context, tx *gorm.DB, user model.RequestUserRegister) (model.User, error)
}
