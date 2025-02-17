package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"s-coder-snippet-sharder/pkg/config"
)

// Connect to postgrest db
func Conn() *pgxpool.Pool {
	connPool, err := pgxpool.New(context.Background(), config.Envs.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to PostgreSQL database: ", err)
	}
	return connPool
}
