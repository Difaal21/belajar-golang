package helpers

import (
	"github.com/Difaal21/belajar-golang/golang-restful-api/model/domain"
	"github.com/Difaal21/belajar-golang/golang-restful-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {

	var categoryResponse []web.CategoryResponse
	for _, category := range categories {
		categoryResponse = append(categoryResponse, ToCategoryResponse(category))
	}
	return categoryResponse
}
