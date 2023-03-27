package services

import (
	"challenges-three/helpers"
	"challenges-three/models"
	"challenges-three/repository"
	"context"
	"log"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type productService struct {
	db          *gorm.DB
	repoProduct repository.ProductRepository
	validator   *validator.Validate
}

func NewProductService(db *gorm.DB, repoProduct repository.ProductRepository, validator_ validator.Validate) *productService {
	return &productService{
		db:          db,
		repoProduct: repoProduct,
		validator:   &validator_,
	}
}

func (service *productService) FindAll(ctx context.Context) ([]models.Product, error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	if products, err := service.repoProduct.FindAll(ctx, tx); err != nil {
		log.Println("Data not found")
		return products, err
	} else {
		// pengganti for loop
		// responseProduct := make([]models.Product, 0, len(products))
		// responseProduct = append(responseProduct, products...)
		return products, nil
	}
}

func (service *productService) FindAllByUserId(ctx context.Context, userID uint) ([]models.Product, error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	products := service.repoProduct.FindAllByUserId(ctx, tx, userID)
	// pengganti for loop
	responseProduct := make([]models.Product, 0, len(products))
	responseProduct = append(responseProduct, products...)
	return responseProduct, nil
}
