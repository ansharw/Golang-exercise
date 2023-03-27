package repository

import (
	"challenges-three/models"
	"context"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type ProductRepositoryMock struct {
	mock.Mock
}

// FindAllByUserId implements ProductRepository
func (rpt *ProductRepositoryMock) FindAllByUserId(ctx context.Context, tx *gorm.DB, userID uint) []models.Product {
	args := rpt.Called(ctx, tx, userID)
	return args.Get(0).([]models.Product)
}

func (rpt *ProductRepositoryMock) FindAll(ctx context.Context, tx *gorm.DB) ([]models.Product, error) {
	args := rpt.Called(ctx, tx)
	return args.Get(0).([]models.Product), args.Error(1)
}
