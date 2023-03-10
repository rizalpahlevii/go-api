package main

import (
	"fmt"
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
	fmt.Println("Server is running on port: ", server.Addr)
	err := server.ListenAndServe()

	helper.PanicIfError(err)
}
