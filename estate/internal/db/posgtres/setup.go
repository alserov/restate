package posgtres

import (
	"context"
	"database/sql"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"time"
)

func MustConnect(dsn string) (*pgxpool.Conn, func()) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		panic("failed to create pool: " + err.Error())
	}

	if err = pool.Ping(ctx); err != nil {
		panic("failed to ping: " + err.Error())
	}

	mustMigrate(stdlib.OpenDBFromPool(pool))

	conn, err := pool.Acquire(ctx)
	if err != nil {
		panic("failed to acquire pool: " + err.Error())
	}

	return conn, func() {
		ctx, cancel = context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err = conn.Conn().Close(ctx); err != nil {
			panic("failed to close connection: " + err.Error())
		}
	}
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
