package repository

import (
	"challenges-three/models"
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	Login(ctx context.Context, tx *gorm.DB, email, pass string) (models.User, error)
	Register(ctx context.Context, tx *gorm.DB, user models.User) (models.User, error)
	FindByEmail(ctx context.Context, tx *gorm.DB, email string) string
}
