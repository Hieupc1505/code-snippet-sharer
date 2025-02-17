package repo

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"s-coder-snippet-sharder/internal/app/account"
	db "s-coder-snippet-sharder/internal/db/sqlc"
)

// create a new account repo
func NewAccountRepo(conn *pgxpool.Pool) account.ReadWriter {
	return db.New(conn)
}
