package db

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Init() *pgxpool.Pool {
	dsn := "postgres://sajha:passwd@localhost:5432/sajha_dev?sslmode=disable"
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatal("Error occured while parsing the connection string")
	}
	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnLifetime = time.Hour

	defaultPgxPool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal("Cannot connect to db")
	}

	return defaultPgxPool
}
