package app

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"pznrestfulapi/helper"
	"time"
)

// Create migration
// migrate create -ext sql -dir db/migrations create_first_table

// Run up migration
// migrate -database "mysql://root@tcp(localhost:3306)/golang_api_database_migration" -path db/migration up

// run down migration
// migrate -database "mysql://root@tcp(localhost:3306)/golang_api_database_migration" -path db/migration down

// root@tcp(localhost:3306)/go_api

func NewDB() *sql.DB {
	connectionString := "root@tcp(localhost:3306)/golang_api"
	db, err := sql.Open("mysql", connectionString)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
