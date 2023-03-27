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
		return products, err
	}
	return products, nil
}

func (repo *productRepository) FindAllByUserId(ctx context.Context, tx *gorm.DB, userID uint) ([]models.Product, error) {
	products := []models.Product{}
	if err := tx.WithContext(ctx).Where("user_id = ?", userID).Order("id DESC").Find(&products).Error; err != nil {
		log.Println("Error finding all products:", err)
		return products, err
	}
	return products, nil
}

// admin
func (repo *productRepository) FindById(ctx context.Context, tx *gorm.DB, id uint) (models.Product, error) {
	product := models.Product{}
	if err := tx.WithContext(ctx).Where("id = ?", id).First(&product).Error; err != nil {
		log.Println("Error finding product:", err)
		return product, err
	}
	return product, nil
}

// user
func (repo *productRepository) FindByUserId(ctx context.Context, tx *gorm.DB, userID uint, id uint) (models.Product, error) {
	product := models.Product{}
	if err := tx.WithContext(ctx).Where("user_id = ? AND id = ?", userID, id).First(&product).Error; err != nil {
		log.Println("Error finding product:", err)
		return product, err
	}
	return product, nil
}

func (repo *productRepository) Create(ctx context.Context, tx *gorm.DB, product models.Product, userID uint) (models.Product, error) {
	product.UserID = userID
	if err := tx.WithContext(ctx).Create(&product).Error; err != nil {
		log.Println("Error creating product:", err)
		return product, err
	}
	return product, nil
}

func (repo *productRepository) Update(ctx context.Context, tx *gorm.DB, product models.Product, id uint) (models.Product, error) {
	if err := tx.WithContext(ctx).Model(&product).Where("id = ?", id).Updates(models.Product{GormModel: models.GormModel{ID: id}, Title: product.Title, Description: product.Description}).Error; err != nil {
		log.Println("Error updating product:", err)
		return product, err
	}
	// to return result after update
	tx.WithContext(ctx).Where("id = ?", id).First(&product)
	return product, nil
}

func (repo *productRepository) Delete(ctx context.Context, tx *gorm.DB, id uint) (models.Product, error) {
	product := models.Product{}
	if err := tx.WithContext(ctx).Where("id = ?", id).First(&product).Error; err != nil {
		log.Println("Error deleting product:", err)
		return product, err
	}
	return product, nil
}
