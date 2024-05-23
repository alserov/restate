package posgtres

import (
	"context"
	"github.com/jackc/pgx/v5"
	"time"
)

func MustConnect(dsn string) (*pgx.Conn, func()) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		panic("failed to connect: " + err.Error())
	}

	if err = conn.Ping(ctx); err != nil {
		panic("failed to ping: " + err.Error())
	}

	//mustMigrate(conn)

	return conn, func() {
		ctx, cancel = context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		_ = conn.Close(ctx)
	}
}
