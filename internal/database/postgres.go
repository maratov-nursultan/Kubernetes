package database

import (
	"database/sql"
	"fmt"
	"github.com/maratov-nursultan/Kubernetes/internal/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"time"
)

func ConnectDatabase(cfg *config.Config) (*bun.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())
	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetMaxOpenConns(90)
	db.SetMaxIdleConns(10)

	err := db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
