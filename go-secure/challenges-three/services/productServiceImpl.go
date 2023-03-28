package services

import (
	"challenges-three/helpers"
	"challenges-three/models"
	"challenges-three/repository"
	"context"
	"log"

	"gorm.io/gorm"
)

type productService struct {
	db          *gorm.DB
	repoProduct repository.ProductRepository
	// validator   *validator.Validate
}

func NewProductService(db *gorm.DB, repoProduct repository.ProductRepository) *productService {
	return &productService{
		db:          db,
		repoProduct: repoProduct,
		// validator:   &validator_,
	}
}

func (service *productService) FindAll(ctx context.Context) ([]models.Product, error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	if products, err := service.repoProduct.FindAll(ctx, tx); err != nil {
		log.Println("Data not found")
		return products, err
	} else {
		return products, nil
	}
}

func (service *productService) FindAllByUserId(ctx context.Context, userID uint) ([]models.Product, error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	if products, err := service.repoProduct.FindAllByUserId(ctx, tx, userID); err != nil {
		log.Println("Data not found")
		return products, err
	} else {
		return products, nil
	}
}

func (service *productService) FindById(ctx context.Context, id uint) (models.Product, error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	if product, err := service.repoProduct.FindById(ctx, tx, id); err != nil {
		log.Println("Data not found")
		return product, err
	} else {
		return product, nil
	}
}

func (service *productService) FindByUserId(ctx context.Context, userID uint, id uint) (models.Product, error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	if product, err := service.repoProduct.FindByUserId(ctx, tx, userID, id); err != nil {
		log.Println("Data not found")
		return product, err
	} else {
		return product, nil
	}
}

func (service *productService) Create(ctx context.Context, product models.Product, userID uint) (models.Product, error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	if product, err := service.repoProduct.Create(ctx, tx, product, userID); err != nil {
		log.Println("Failed to create product")
		return product, err
	} else {
		return product, nil
	}
}

func (service *productService) Update(ctx context.Context, product models.Product, id uint) (models.Product, error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	if product, err := service.repoProduct.Update(ctx, tx, product, id); err != nil {
		log.Println("Failed to update product")
		return product, err
	} else {
		return product, nil
	}
}

func (service *productService) Delete(ctx context.Context, id uint) (models.Product, error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	if product, err := service.repoProduct.Delete(ctx, tx, id); err != nil {
		log.Println("Failed to delete product")
		return product, err
	} else {
		return product, nil
	}
}
