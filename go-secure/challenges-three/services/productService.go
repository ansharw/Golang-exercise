package services

import (
	"challenges-three/models"
	"context"
)

type ProductService interface {
	FindAll(ctx context.Context) ([]models.Product, error)
	// FindById(ctx context.Context, id int) (response.ResponseUserContentByContentId, error)
	// FindByStatusPublish(ctx context.Context) ([]response.ResponseUserContentByContentId, error)
	// FindByStatusDraft(ctx context.Context) ([]response.ResponseUserContentByContentId, error)
	// Create(ctx context.Context, request request.RequestCreateContent, user_id int) (response.ResponseUserContentCreated, error)
	// Update(ctx context.Context, request request.RequestUpdateContent, user_id, content_id int) (response.ResponseUserContentUpdateByContentId, error)
	// Delete(ctx context.Context, user_id, content_id int) error
	FindAllByUserId(ctx context.Context, userID uint) ([]models.Product, error)
}
