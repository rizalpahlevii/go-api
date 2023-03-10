package app

import (
	"github.com/julienschmidt/httprouter"
	"pznrestfulapi/controller"
	"pznrestfulapi/exception"
)

func NewRouter(
	categoryController controller.CategoryController,
	authController controller.AuthController,
) *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.POST("/api/auth/login", authController.Login)
	router.POST("/api/auth/logout", authController.Logout)
	router.PanicHandler = exception.ErrorHandler
	return router
}
