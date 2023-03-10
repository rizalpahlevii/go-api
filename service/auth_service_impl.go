package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"pznrestfulapi/exception"
	"pznrestfulapi/helper"
	"pznrestfulapi/model/web"
	"pznrestfulapi/repository"
	"strconv"
)

type AuthServiceImpl struct {
	authRepository repository.AuthRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func (service *AuthServiceImpl) Login(ctx context.Context, request web.LoginRequest) web.LoginResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.authRepository.Login(ctx, tx, request.Username, request.Password)
	// Append user token to response
	user.Token, _ = GenerateToken(strconv.Itoa(user.Id))
	if err != nil {
		panic(exception.NewLoginError(err.Error()))
	}

	return helper.ToLoginResponse(user)

}

func NewAuthService(authRepository repository.AuthRepository, DB *sql.DB, validate *validator.Validate) *AuthServiceImpl {
	return &AuthServiceImpl{authRepository: authRepository, DB: DB, Validate: validate}
}
