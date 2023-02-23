package main

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"pznrestfulapi/app"
	"pznrestfulapi/controller"
	"pznrestfulapi/helper"
	"pznrestfulapi/middleware"
	"pznrestfulapi/repository"
	"pznrestfulapi/service"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
