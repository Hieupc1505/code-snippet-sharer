package snippet

import "s-coder-snippet-sharder/internal/app"

type Lang string

func (l Lang) String() string { return string(l) }

func NewLang(lang string) (Lang, error) {
	switch lang {
	case "yaml", "js", "go", "py", "java":
		return Lang(lang), nil
	default:
		return "", app.ErrUnsupportedLanguage
	}
}
