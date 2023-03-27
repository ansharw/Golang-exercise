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

func (rpt *ProductRepositoryMock) FindAll(ctx context.Context, tx *gorm.DB) ([]models.Product, error) {
	args := rpt.Called(ctx, tx)
	return args.Get(0).([]models.Product), args.Error(1)
}

func (rpt *ProductRepositoryMock) FindAllByUserId(ctx context.Context, tx *gorm.DB, userID uint) ([]models.Product, error) {
	args := rpt.Called(ctx, tx, userID)
	return args.Get(0).([]models.Product), args.Error(1)
}

func (rpt *ProductRepositoryMock) FindById(ctx context.Context, tx *gorm.DB, id uint) (models.Product, error) {
	args := rpt.Called(ctx, tx, id)
	return args.Get(0).(models.Product), args.Error(1)
}

func (rpt *ProductRepositoryMock) FindByUserId(ctx context.Context, tx *gorm.DB, userID uint, id uint) (models.Product, error) {
	args := rpt.Called(ctx, tx, userID, id)
	return args.Get(0).(models.Product), args.Error(1)
}

func (rpt *ProductRepositoryMock) Create(ctx context.Context, tx *gorm.DB, userID uint) (models.Product, error) {
	args := rpt.Called(ctx, tx, userID)
	return args.Get(0).(models.Product), args.Error(1)
}

func (rpt *ProductRepositoryMock) Update(ctx context.Context, tx *gorm.DB, id uint) (models.Product, error) {
	args := rpt.Called(ctx, tx, id)
	return args.Get(0).(models.Product), args.Error(1)
}

func (rpt *ProductRepositoryMock) Delete(ctx context.Context, tx *gorm.DB, id uint) (models.Product, error) {
	args := rpt.Called(ctx, tx, id)
	return args.Get(0).(models.Product), args.Error(1)
}
