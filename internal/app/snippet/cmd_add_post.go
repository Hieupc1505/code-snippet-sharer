package snippet

import (
	"context"
	"fmt"
	"github.com/lithammer/shortuuid"
	"s-coder-snippet-sharder/internal/app"
	db "s-coder-snippet-sharder/internal/db/sqlc"
	"s-coder-snippet-sharder/pkg/errsx"
)

func NewCodeSnippet(lang, title, snippet string, public bool) (*db.AddParams, error) {
	var errs errsx.Map
	langType, err := NewLang(lang)
	if err != nil {
		errs.Set("lang", err)
	}

	titlePost, err := NewTitle(title)
	if err != nil {
		errs.Set("title", err)
	}

	snippetContent, err := NewSnippet(snippet)
	if err != nil {
		errs.Set("snippet", err)
	}

	if errs != nil {
		return nil, fmt.Errorf("%w: %v", app.ErrInvalidInput, errs)
	}

	post := NewSnippetParam(langType, titlePost, snippetContent, public)
	return post, nil

}

func (s *Service) CreateNewCodePost(ctx context.Context, lang, title, code string) (*db.AddParams, error) {
	post, err := NewCodeSnippet(lang, title, code, true)
	if err != nil {
		return nil, err
	}
	if err := s.repo.Add(ctx, db.AddParams{
		Lang:    lang,
		Title:   title,
		Snippet: code,
		Slug:    shortuuid.New(),
		Public:  true,
	}); err != nil {
		return nil, err
	}
	return post, nil
}
