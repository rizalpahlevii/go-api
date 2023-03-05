package main

import (
	"net/http"
	"pznrestfulapi/helper"
	"pznrestfulapi/middleware"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:8080",
		Handler: authMiddleware,
	}
}

func main() {

	server := InitializedServer()

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
