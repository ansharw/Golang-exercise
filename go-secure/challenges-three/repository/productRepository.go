package repository

import (
	"challenges-three/models"
	"context"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll(ctx context.Context, tx *gorm.DB) ([]models.Product, error)
	FindAllByUserId(ctx context.Context, tx *gorm.DB, userID uint) ([]models.Product, error)
	FindById(ctx context.Context, tx *gorm.DB, id uint) (models.Product, error)
	FindByUserId(ctx context.Context, tx *gorm.DB, userID uint, id uint) (models.Product, error)
	Create(ctx context.Context, tx *gorm.DB, product models.Product, userID uint) (models.Product, error)
	Update(ctx context.Context, tx *gorm.DB, product models.Product, id uint) (models.Product, error)
	Delete(ctx context.Context, tx *gorm.DB, id uint) (models.Product, error)
}
