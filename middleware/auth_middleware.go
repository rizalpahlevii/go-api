package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"pznrestfulapi/helper"
	"pznrestfulapi/model/web"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	tokenString := request.Header.Get("Authorization")

	if request.URL.Path == "/api/auth/login" {
		middleware.Handler.ServeHTTP(writer, request)
		return
	}

	if tokenString == "" {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}
		return []byte("secret"), nil
	})

	if err != nil || !token.Valid {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}
	middleware.Handler.ServeHTTP(writer, request)
}

//func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
//	if "RAHASIA" == request.Header.Get("X-API-Key") {
//		middleware.Handler.ServeHTTP(writer, request)
//	} else {
//		writer.Header().Set("Content-Type", "application/json")
//		writer.WriteHeader(http.StatusUnauthorized)
//		webResponse := web.WebResponse{
//			Code:   http.StatusUnauthorized,
//			Status: "Unauthorized",
//		}
//		helper.WriteToResponseBody(writer, webResponse)
//	}
//}
