// Package Repository itu berisi tentang jembatan ke database
package repository

import "github.com/Difaal21/belajar-golang/unit-test-mock/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
