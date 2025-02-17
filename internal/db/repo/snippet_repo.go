package repo

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"s-coder-snippet-sharder/internal/app/snippet"
	db "s-coder-snippet-sharder/internal/db/sqlc"
)

func NewSnippetRepo(conn *pgxpool.Pool) snippet.ReadWriter {
	return db.New(conn)
}
