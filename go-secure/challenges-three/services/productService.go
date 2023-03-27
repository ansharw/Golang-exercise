package services

import (
	"challenges-three/models"
	"context"
)

type ProductService interface {
	FindAll(ctx context.Context) ([]models.Product, error)
	FindAllByUserId(ctx context.Context, userID uint) ([]models.Product, error)
	FindById(ctx context.Context, id uint) (models.Product, error)
	FindByUserId(ctx context.Context, userID uint, id uint) (models.Product, error)
	Create(ctx context.Context, product models.Product, userID uint) (models.Product, error)
	Update(ctx context.Context, product models.Product, id uint) (models.Product, error)
	Delete(ctx context.Context, id uint) (models.Product, error)
}
