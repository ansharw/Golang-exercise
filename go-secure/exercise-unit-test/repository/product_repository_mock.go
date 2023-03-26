package repository

import (
	"exercise-unit-test/entity"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (RPM *ProductRepositoryMock) FindById(id string) *entity.Product {
	args := RPM.Mock.Called(id)

	if args.Get(0) == nil {
		return nil
	}

	product := args.Get(0).(entity.Product)

	return &product
}