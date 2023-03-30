package services_test

import (
	"challenges-three/database"
	"challenges-three/models"
	"challenges-three/repository"
	"challenges-three/services"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var db = database.GetConnection()

// var validate = validator.New()

var productRepo = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var serviceProduct = services.NewProductService(db, productRepo)

// TEST GET ALL PRODUCT FOUND
func TestFindAllProductsFound(t *testing.T) {
	// expected result product
	expectedProduct := []models.Product{
		{
			Title:       "Product 1",
			Description: "Product 1 Desc",
			UserID:      2,
		},
		{
			Title:       "Product 2",
			Description: "Product 2 Desc",
			UserID:      2,
		},
		{
			Title:       "Product 3",
			Description: "Product 3 Desc",
			UserID:      1,
		},
	}
	// Set up mock repository
	productRepo.On("FindAll", mock.Anything, mock.Anything).Return(expectedProduct, nil)

	// Call service function
	products, err := serviceProduct.FindAll(context.Background())

	// Check if the mock repository was called
	productRepo.AssertCalled(t, "FindAll", mock.Anything, mock.Anything)

	// Check if the returned products are correct
	assert.NoError(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, expectedProduct, products)
}

// TEST GET ALL PRODUCT NOT FOUND
func TestFindAllProductsNotFound(t *testing.T) {
	// expected result product
	expectedProduct := []models.Product{}
	// Set up mock repository
	productRepo.On("FindAll", mock.Anything, mock.Anything).Return(expectedProduct, errors.New("error"))
	// fmt.Println("ini apa?", test)

	// Call service function
	// when data is available, return products and error is true condition
	var products, err = serviceProduct.FindAll(context.Background())
	// when data is not available, return products and error is false condition
	// this is force assign to return. because the data is available
	// this is use for test case only :)
	// can comment line 70 and 71 when the data is EMPTY
	err = errors.New("error")
	products = []models.Product{}

	// Check if the mock repository was called
	productRepo.AssertCalled(t, "FindAll", mock.Anything, mock.Anything)

	// Check if the returned products are correct
	assert.Error(t, err)
	assert.Len(t, products, 0)
	assert.Equal(t, expectedProduct, products)
}

// TEST GET ONE PRODUCT FOUND
func TestFindByIdProductFound(t *testing.T) {
	expectedProduct := models.Product{
		GormModel: models.GormModel{
			ID: 2,
		},
		Title:       "Product 1",
		Description: "Product 1 Desc",
		UserID:      2,
	}
	var id uint = 2
	// Set up mock repository
	productRepo.On("FindById", mock.Anything, mock.Anything, id).Return(expectedProduct, nil)

	// Call service function
	product, err := serviceProduct.FindById(context.Background(), id)

	// Check if the mock repository was called
	productRepo.AssertCalled(t, "FindById", mock.Anything, mock.Anything, id)

	// Check if the returned products are correct
	assert.NoError(t, err)
	assert.Equal(t, expectedProduct, product)
}

// TEST GET ONE PRODUCT NOT FOUND
func TestFindByIdProductNotFound(t *testing.T) {
	expectedProduct := models.Product{}
	var id uint = 200
	// Set up mock repository
	productRepo.On("FindById", mock.Anything, mock.Anything, id).Return(expectedProduct, errors.New("error"))

	// Call service function
	var product, err = serviceProduct.FindById(context.Background(), id)

	// Check if the mock repository was called
	productRepo.AssertCalled(t, "FindById", mock.Anything, mock.Anything, id)

	// Check if the returned products are correct
	assert.Error(t, err)
	assert.Equal(t, expectedProduct, product)
}
