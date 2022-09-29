//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	"github.com/Difaal21/belajar-golang/golang-restful-api/app"
	"github.com/Difaal21/belajar-golang/golang-restful-api/controller"
	"github.com/Difaal21/belajar-golang/golang-restful-api/middleware"
	"github.com/Difaal21/belajar-golang/golang-restful-api/repository"
	"github.com/Difaal21/belajar-golang/golang-restful-api/services"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

// Grouping
var categorySet = wire.NewSet(
	/* Dibinding karena balikan NewCategoryRepositoryImpl itu sebuah struct */
	repository.NewCategoryRepositoryImpl,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),

	/* Dibinding karena balikan NewCategoryServiceImpl itu sebuah struct */
	services.NewCategoryServiceImpl,
	wire.Bind(new(services.CategoryService), new(*services.CategoryServiceImpl)),

	/* Dibinding karena balikan NewCategoryControllerImpl itu sebuah struct */
	controller.NewCategoryControllerImpl,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func InitializedServer() *http.Server {

	wire.Build(
		app.NewDB,
		validator.New,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
