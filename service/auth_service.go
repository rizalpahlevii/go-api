package service

import (
	"context"
	"pznrestfulapi/model/web"
)

type AuthService interface {
	Login(ctx context.Context, request web.LoginRequest) web.LoginResponse
}
