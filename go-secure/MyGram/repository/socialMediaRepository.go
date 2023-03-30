package repository

import (
	"MyGram/model"
	"context"

	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	FindAllByUserId(ctx context.Context, tx *gorm.DB, userID uint) ([]model.SocialMedia, error)
	FindByUserId(ctx context.Context, tx *gorm.DB, userID uint, id uint) (model.SocialMedia, error)
	Create(ctx context.Context, tx *gorm.DB, req model.RequestSocialMedia, userID uint) (model.SocialMedia, error)
	Update(ctx context.Context, tx *gorm.DB, req model.RequestSocialMedia, id, userID uint) (model.SocialMedia, error)
	Delete(ctx context.Context, tx *gorm.DB, id, userID uint) (model.SocialMedia, error)
}
