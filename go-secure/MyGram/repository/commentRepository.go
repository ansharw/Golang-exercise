package repository

import (
	"MyGram/model"
	"context"

	"gorm.io/gorm"
)

type CommentRepository interface {
	FindAllByPhotoId(ctx context.Context, tx *gorm.DB, photoID uint) ([]model.Comments, error)
	FindByPhotoId(ctx context.Context, tx *gorm.DB, photoID uint, id uint) (model.Comments, error)
	Create(ctx context.Context, tx *gorm.DB, req model.RequestComments, userID, photoID uint) (model.Comments, error)
	Update(ctx context.Context, tx *gorm.DB, req model.RequestComments, id, userID, photoID uint) (model.Comments, error)
	Delete(ctx context.Context, tx *gorm.DB, id, userID, photoID uint) error
}
