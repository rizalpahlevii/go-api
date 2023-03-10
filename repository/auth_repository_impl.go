package repository

import (
	"context"
	"database/sql"
	"errors"
	"pznrestfulapi/helper"
	"pznrestfulapi/model/domain"
)

type AuthRepositoryImpl struct {
}

func NewAuthRepository() *AuthRepositoryImpl {
	return &AuthRepositoryImpl{}
}

func (a AuthRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, username string, password string) (domain.User, error) {
	SQL := "SELECT id, username, password FROM users WHERE username = ?"
	rows, err := tx.QueryContext(ctx, SQL, username)
	helper.PanicIfError(err)
	defer rows.Close()

	if !rows.Next() {
		return domain.User{}, errors.New("user not found")
	}

	// Compare password
	user := domain.User{}
	err = rows.Scan(&user.Id, &user.Username, &user.Password)
	hashPassword := user.Password
	helper.PanicIfError(err)
	if !helper.CheckPasswordHash(password, hashPassword) {
		return domain.User{}, errors.New("wrong password")
	}
	return user, nil
}
