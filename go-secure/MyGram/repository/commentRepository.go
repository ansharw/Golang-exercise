package repository

import (
	"MyGram/model"
	"context"

	"gorm.io/gorm"
)

type CommentRepository interface {
	FindAllByPhotoId(ctx context.Context, tx *gorm.DB, photoID uint) ([]model.Comment, error)
	FindByPhotoId(ctx context.Context, tx *gorm.DB, photoID uint, id uint) (model.Comment, error)
	Create(ctx context.Context, tx *gorm.DB, comment model.Comment, userID, photoID uint) (model.Comment, error)
	Update(ctx context.Context, tx *gorm.DB, comment model.Comment, id, userID, photoID uint) (model.Comment, error)
	Delete(ctx context.Context, tx *gorm.DB, id, userID, photoID uint) (model.Comment, error)
}
