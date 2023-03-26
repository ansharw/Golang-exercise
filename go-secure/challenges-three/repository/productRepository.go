package repository

import (
	"challenges-three/models"
	"context"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll(ctx context.Context, tx *gorm.DB) []models.Product
	// FindById(ctx context.Context, tx *sql.Tx, id int) domain.Content
	// Create(ctx context.Context, tx *sql.Tx, content domain.Content) *int
	// Update(ctx context.Context, tx *sql.Tx, content domain.Content)
	// Delete(ctx context.Context, tx *sql.Tx, id int)
	FindAllByUserId(ctx context.Context, tx *gorm.DB, userID uint) []models.Product
}
