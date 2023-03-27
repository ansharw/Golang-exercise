package repository

import (
	"challenges-three/models"
	"context"
	"log"

	"gorm.io/gorm"
)

type productRepository struct {
}

func NewProductRepository() *productRepository {
	return &productRepository{}
}

func (repo *productRepository) FindAll(ctx context.Context, tx *gorm.DB) ([]models.Product, error) {
	products := []models.Product{}
	if err := tx.WithContext(ctx).Order("id DESC").Find(&products).Error; err != nil {
		log.Println("Error finding all products:", err)
		// default value []struct is nil if data not found
		return products, nil
	}
	return products, nil
}

func (repo *productRepository) FindAllByUserId(ctx context.Context, tx *gorm.DB, userID uint) []models.Product {
	products := []models.Product{}
	if err := tx.WithContext(ctx).Where("user_id = ?", userID).Order("id DESC").Find(&products).Error; err != nil {
		log.Println("Error finding all products:", err)
	}
	return products
}

func (repo *productRepository) FindById(ctx context.Context, tx *gorm.DB, userID uint, id int) models.Product {
	product := models.Product{}
	if err := tx.WithContext(ctx).Where("user_id = ? AND id = ?", userID, id).Order("id DESC").First(&product).Error; err != nil {
		log.Println("Error finding product:", err)
	}
	return product
}
