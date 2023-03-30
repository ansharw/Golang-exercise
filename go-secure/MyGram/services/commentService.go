package services

import (
	"MyGram/model"
	"context"
)

type CommentService interface {
	FindAllByPhotoId(ctx context.Context, photoID uint) ([]model.Comment, error)
	FindByPhotoId(ctx context.Context, photoID uint, id uint) (model.Comment, error)
	Create(ctx context.Context, req model.RequestComment, userID, photoID uint) (model.Comment, error)
	Update(ctx context.Context, req model.RequestComment, id, userID, photoID uint) (model.Comment, error)
	Delete(ctx context.Context, id, userID, photoID uint) (model.Comment, error)
}
