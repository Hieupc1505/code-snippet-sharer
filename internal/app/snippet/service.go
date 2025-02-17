package snippet

import (
	"context"
	db "s-coder-snippet-sharder/internal/db/sqlc"
)

type Reader interface {
	GetBySlug(ctx context.Context, slug string) (db.CodeSnippet, error)
	GetPublicSnippets(ctx context.Context) ([]db.CodeSnippet, error)
	GetRecentPosts(ctx context.Context) ([]db.CodeSnippet, error)
}

type Writer interface {
	Add(ctx context.Context, arg db.AddParams) error
	UpdateViewCount(ctx context.Context, slug string) error
}

type ReadWriter interface {
	Reader
	Writer
}

type Service struct {
	repo ReadWriter
	Db   db.DBTX
}

func NewService(ctx context.Context, repo ReadWriter) (*Service, error) {
	return &Service{repo: repo}, nil
}
