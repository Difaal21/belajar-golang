package main

import (
	"net/http"

	"github.com/Difaal21/belajar-golang/golang-restful-api/app"
	"github.com/Difaal21/belajar-golang/golang-restful-api/controller"
	"github.com/Difaal21/belajar-golang/golang-restful-api/exception"
	"github.com/Difaal21/belajar-golang/golang-restful-api/helpers"
	"github.com/Difaal21/belajar-golang/golang-restful-api/middleware"
	"github.com/Difaal21/belajar-golang/golang-restful-api/repository"
	"github.com/Difaal21/belajar-golang/golang-restful-api/services"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := services.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/v1/categories", categoryController.FindAll)
	router.GET("/api/v1/categories/:categoryId", categoryController.FindById)
	router.POST("/api/v1/categories", categoryController.Create)
	router.PUT("/api/v1/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/v1/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helpers.PanicIfError(err)
}
