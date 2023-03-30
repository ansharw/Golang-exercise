package repository

import (
	"MyGram/model"
	"context"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	FindAllByUserId(ctx context.Context, tx *gorm.DB, userID uint) ([]model.Photo, error)
	FindByUserId(ctx context.Context, tx *gorm.DB, userID uint, id uint) (model.Photo, error)
	Create(ctx context.Context, tx *gorm.DB, photo model.Photo, userID uint) (model.Photo, error)
	Update(ctx context.Context, tx *gorm.DB, photo model.Photo, id, userID uint) (model.Photo, error)
	Delete(ctx context.Context, tx *gorm.DB, id, userID uint) (model.Photo, error)
}
