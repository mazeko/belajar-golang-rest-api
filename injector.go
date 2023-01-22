//go:build wireinject
// +build wireinject

package main

import (
	"belajar-goalng-rest-api/app"
	"belajar-goalng-rest-api/controller"
	"belajar-goalng-rest-api/middleware"
	"belajar-goalng-rest-api/repository"
	"belajar-goalng-rest-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepo), new(*repository.CategoryRepoImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func InitialServer() *http.Server {
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
