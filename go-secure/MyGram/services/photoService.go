package services

import (
	"MyGram/model"
	"context"
)

type PhotoService interface {
	FindAllByUserId(ctx context.Context, userID uint) ([]model.Photo, error)
	FindByUserId(ctx context.Context, userID uint, id uint) (model.Photo, error)
	Create(ctx context.Context, photo model.Photo, userID uint) (model.Photo, error)
	Update(ctx context.Context, photo model.Photo, id, userID uint) (model.Photo, error)
	Delete(ctx context.Context, id, userID uint) (model.Photo, error)
}
