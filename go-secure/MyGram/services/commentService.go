package services

import (
	"MyGram/model"
	"context"
)

type CommentService interface {
	FindAllByPhotoId(ctx context.Context, photoID uint) ([]model.Comments, error)
	FindByPhotoId(ctx context.Context, photoID uint, id uint) (model.Comments, error)
	Create(ctx context.Context, req model.RequestComments, userID, photoID uint) (model.Comments, error)
	Update(ctx context.Context, req model.RequestComments, id, userID, photoID uint) (model.Comments, error)
	Delete(ctx context.Context, id, userID, photoID uint) error
}
