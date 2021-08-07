package sql

import (
	"database/sql"
	"fmt"

	"github.com/asenlog/hexagonal-architecture/config"
)

// New creates a handle to a Database.
func New(cfg config.DB, host string) (*sql.DB, error) {
	db, err := sql.Open("mysql",
		fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True", cfg.Username, cfg.Password, host, cfg.Port, cfg.DB))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("cannot ping DB : %w", err)
	}

	return db, nil
}
