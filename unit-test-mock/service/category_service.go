// Package Service itu berisi Business Logic
package service

import (
	"errors"
	"unit-test-mock/entity"
	"unit-test-mock/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("CATEGORY NOT FOUND")
	} else {
		return category, nil
	}
}
