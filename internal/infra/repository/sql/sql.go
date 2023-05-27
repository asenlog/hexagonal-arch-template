package sql

import (
	"app/internal/config"
	"context"
	"fmt"
	"net/url"

	// PG driver.
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

// New creates a handle to a Database.
func New(ctx context.Context, cfg config.DB, host string) (*sqlx.DB, error) {
	opt := url.Values{}
	opt.Add("sslmode", cfg.SSLMode)
	opt.Add("timezone", "UTC")

	con := &url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(cfg.Username, cfg.Password),
		Host:     host,
		Path:     cfg.DB,
		RawQuery: opt.Encode(),
	}

	// underneath uses sql.Open that is concurrent safe.
	db, err := sqlx.Connect("pgx", con.String())
	if err != nil {
		return nil, fmt.Errorf("sql.Open failed: %w", err)
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("db.Ping failed: %w", err)
	}

	return db, nil
}
