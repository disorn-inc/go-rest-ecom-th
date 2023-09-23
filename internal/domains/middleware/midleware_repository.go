package middleware

import "github.com/jmoiron/sqlx"

type IMiddlewareRepository interface {}

type MiddlewareRepository struct {
	db *sqlx.DB
}

func NewMiddlewareRepository(db *sqlx.DB) IMiddlewareRepository {
	return &MiddlewareRepository{
		db: db,
	}
}