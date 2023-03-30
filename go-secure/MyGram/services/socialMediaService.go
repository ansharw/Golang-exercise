package services

import (
	"MyGram/model"
	"context"
)

type SocialMediaService interface {
	FindAllByUserId(ctx context.Context, userID uint) ([]model.SocialMedia, error)
	FindByUserId(ctx context.Context, userID uint, id uint) (model.SocialMedia, error)
	Create(ctx context.Context, req model.RequestSocialMedia, userID uint) (model.SocialMedia, error)
	Update(ctx context.Context, req model.RequestSocialMedia, id, userID uint) (model.SocialMedia, error)
	Delete(ctx context.Context, id, userID uint) (model.SocialMedia, error)
}
