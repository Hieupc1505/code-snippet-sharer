package snippet

import (
	"fmt"
	db "s-coder-snippet-sharder/internal/db/sqlc"
)

func NewSnippetParam(lang Lang, title Title, snippet Snippet, public bool) *db.AddParams {
	return &db.AddParams{
		Slug:    fmt.Sprintf("%s_%s", lang, title),
		Title:   title.String(),
		Snippet: snippet.String(),
		Lang:    lang.String(),
		Public:  public,
	}
}
