package services

import (
	"exercise-unit-test/entity"
	"exercise-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepo = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{Repository: productRepo}

func TestProductServiceGetOneProductNotFound(t *testing.T)  {
	productRepo.Mock.On("FindById", "1").Return(nil)

	product, err := productService.GetOneProduct("1")

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), "error to response has to be 'product not found'")
}

func TestProductServiceGetOneProduct(t *testing.T)  {
	product := entity.Product{
		Id:   "2",
		Name: "Kaca mata",
	}

	productRepo.Mock.On("FindById", "2").Return(product)

	result, err := productService.GetOneProduct("2")

	assert.Nil(t, err)

	assert.NotNil(t, result)

	assert.Equal(t, product.Id, result.Id, "result has to be '2'")
	assert.Equal(t, product.Name, result.Name, "result has to be 'Kaca mata'")
	assert.Equal(t, &product, result, "result has to be a product data with id '2'")
}