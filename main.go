package main

import (
	"belajar-goalng-rest-api/app"
	"belajar-goalng-rest-api/controller"
	"belajar-goalng-rest-api/exception"
	"belajar-goalng-rest-api/helper"
	"belajar-goalng-rest-api/middleware"
	"belajar-goalng-rest-api/repository"
	"belajar-goalng-rest-api/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categtoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categtoryController.FindAll)
	router.GET("/api/categories/:categoryId", categtoryController.FindById)
	router.POST("/api/categories", categtoryController.Create)
	router.PUT("/api/categories/:categoryId", categtoryController.Update)
	router.DELETE("/api/categories/:categoryId", categtoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicError(err)
}
