package databases

import (
	"log/slog"

	"github.com/disorn-inc/go-rest-ecom-th/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func DbConnect(cfg config.IDbConfig) *sqlx.DB {
	// slog.Info("Connecting to database...", "url", cfg.Url())
	db, err := sqlx.Connect("pgx", cfg.Url())
	if err != nil {
		slog.Error("Failed to connect to database", "err", err.Error())
		panic(err)
	}
	db.SetMaxOpenConns(cfg.MaxConnections())
	return db
}