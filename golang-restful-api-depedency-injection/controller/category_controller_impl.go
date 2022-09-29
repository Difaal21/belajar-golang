package controller

import (
	"net/http"
	"strconv"

	"github.com/Difaal21/belajar-golang/golang-restful-api/helpers"
	"github.com/Difaal21/belajar-golang/golang-restful-api/model/web"
	"github.com/Difaal21/belajar-golang/golang-restful-api/services"
	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService services.CategoryService
}

func NewCategoryControllerImpl(categoryService services.CategoryService) *CategoryControllerImpl {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	categoryCreateRequest := web.CategoryCreateRequest{}
	helpers.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helpers.WriteToResponseBody(writer, webResponse)

}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helpers.ReadFromRequestBody(request, &categoryUpdateRequest)

	categoryId := params.ByName("categoryId")
	// Konversi ke string
	id, err := strconv.Atoi(categoryId)
	helpers.PanicIfError(err)

	categoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helpers.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	// Konversi ke string
	id, err := strconv.Atoi(categoryId)
	helpers.PanicIfError(err)

	controller.CategoryService.Delete(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helpers.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helpers.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helpers.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	categoryResponses := controller.CategoryService.FindAll(request.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}

	helpers.WriteToResponseBody(writer, webResponse)
}
