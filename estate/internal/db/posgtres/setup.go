package posgtres

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func MustConnect(dsn string) *sqlx.DB {
	conn, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		panic("failed to connect: " + err.Error())
	}

	if err = conn.Ping(); err != nil {
		panic("failed to ping: " + err.Error())
	}

	mustMigrate(conn.DB)

	return conn
}

const (
	migrationsDir = "./internal/db/migrations"
)

func mustMigrate(conn *sql.DB) {
	if err := goose.SetDialect("postgres"); err != nil {
		panic("failed to set dialect: " + err.Error())
	}

	if err := goose.Up(conn, migrationsDir); err != nil {
		panic(err)
	}
}
