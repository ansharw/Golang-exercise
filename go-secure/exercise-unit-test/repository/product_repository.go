package repository

import "exercise-unit-test/entity"

type ProductRepository interface {
	FindById(id string) *entity.Product
}