// Package Service itu berisi Business Logic
package service

import (
	"errors"

	"github.com/Difaal21/belajar-golang/unit-test-mock/repository"

	"github.com/Difaal21/belajar-golang/unit-test-mock/entity"
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
