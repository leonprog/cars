package db

import (
	"api/internal/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

func New(ctx context.Context, c *config.Config) *pgx.Conn {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.DbConfig.Username, c.DbConfig.Password, c.DbConfig.Host, c.DbConfig.Port, c.DbConfig.Db)
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		panic(err)
	}

	return conn
}
