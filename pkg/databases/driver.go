package databases

import "github.com/jmoiron/sqlx"

type Driver interface {
	GetPostgres() *sqlx.DB
}

type drivers struct {
	postgres *sqlx.DB
}

func NewDrivers(
	postgres *sqlx.DB,
) Driver {
	return &drivers{
		postgres: postgres,
	}
}

func (d *drivers) GetPostgres() *sqlx.DB {
	return d.postgres
}