// Package Service itu berisi Business Logic
package service

import (
	"testing"
	"unit-test-mock/entity"
	"unit-test-mock/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}

var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {

	// Ketika function FindById dipanggil dengan parameter 1 return nil
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	// Ambil data
	category, err := categoryService.Get("1")

	// Kalau category nil
	assert.Nil(t, category)

	// Kalau category tidal nil
	assert.NotNil(t, err)
}

func TestCategoryService_GetFound(t *testing.T) {
	category := entity.Category{
		Id:   "1",
		Name: "laptop",
	}

	categoryRepository.Mock.On("FindById", "2").Return(category)

	result, error := categoryService.Get("2")

	assert.Nil(t, error)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}
