package repository

import (
	"context"
	"database/sql"
	"pznrestfulapi/model/domain"
)

type AuthRepository interface {
	Login(ctx context.Context, tx *sql.Tx, username string, password string) (domain.User, error)
}
