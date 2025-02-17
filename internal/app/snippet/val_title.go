package snippet

import (
	"s-coder-snippet-sharder/internal/app"
	"strings"
)

const (
	minTitleLen = 10
	maxTitleLen = 100
)

type Title string

func (a Title) String() string { return string(a) }

func NewTitle(title string) (Title, error) {
	title = strings.TrimSpace(title)
	if len(title) < minTitleLen || len(title) > maxTitleLen {
		return "", app.ErrEmailInvalidLength
	}

	return Title(title), nil

}
