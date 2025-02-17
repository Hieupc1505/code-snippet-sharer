package snippet

import (
	"context"
	"errors"
	db "s-coder-snippet-sharder/internal/db/sqlc"
)

var (
	ErrSlugEmpty = errors.New("slug is empty")
)

func (s *Service) GetPublicSnippets(ctx context.Context) ([]db.CodeSnippet, error) {
	return s.repo.GetRecentPosts(ctx)
}

func (s *Service) GetRecentSnippets(ctx context.Context) ([]db.CodeSnippet, error) {
	return s.repo.GetRecentPosts(ctx)
}

func (s *Service) GetSnippetBySlug(ctx context.Context, slug string) (*db.CodeSnippet, error) {

	snippet, err := s.repo.GetBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}

	return &snippet, nil
}
