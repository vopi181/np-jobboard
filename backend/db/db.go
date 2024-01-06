package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"go4.org/syncutil"
)

func Get(ctx context.Context) (*pgxpool.Pool, error) {
	err := once.Do(func() error {
		var err error
		pool, err = setup(ctx)
		return err
	})

	return pool, err
}

var (
	once syncutil.Once
	pool *pgxpool.Pool
)

var secrets struct {
	ExternalDBPassword string
}

func setup(ctx context.Context) (*pgxpool.Pool, error) {
	connString := fmt.Sprintf("postgresql://%s:%s@pgmain.c3pfz1bc4dbi.us-east-1.rds.amazonaws.com",
		"postgres", secrets.ExternalDBPassword)
	return pgxpool.Connect(ctx, connString)
}
