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

func (repo *productRepository) FindAll(ctx context.Context, tx *gorm.DB) []models.Product {
	products := []models.Product{}
	// if res := tx.Order("id DESC").Find(&products); res.RowsAffected == 0 {
	// 	log.Println("No data product")
	// }
	// return products

	if err := tx.WithContext(ctx).Order("id DESC").Find(&products).Error; err != nil {
		log.Println("Error finding all products:", err)
	}
	return products
}

func (repo *productRepository) FindAllByUserId(ctx context.Context, tx *gorm.DB, userID uint) []models.Product {
	products := []models.Product{}
	// if res := tx.Order("id DESC").Find(&products); res.RowsAffected == 0 {
	// 	log.Println("No data product")
	// }
	// return products

	if err := tx.WithContext(ctx).Where("user_id = ?", userID).Order("id DESC").Find(&products).Error; err != nil {
		log.Println("Error finding all products:", err)
	}
	return products
}