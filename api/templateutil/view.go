package templateutil

import (
	"fmt"
	"html/template"
	"log/slog"
	"path/filepath"
)

const TemplateDir = "/home/hieupc05/dev/s-coder-snippet-sharder/api/web/views" // path to views source directory

type View struct {
	NamedEndpoints NamedEndpoints
	AppName        string
	AppTitle       string
}

func ParseFiles() (*template.Template, error) {
	pattern := filepath.Join(TemplateDir, "*", "com_*.tmpl")
	tmpl := template.New("")
	tmpl.Funcs(template.FuncMap{
		"GetFromMap":  GetFromMap,
		"Map":         TmplMap,
		"Slice":       TmplSlice,
		"WithComData": WithComData,
		"substr":      TmplSubstr,
		//"NoEscape":    NoEscape,
	})

	var err error
	tmpl, err = tmpl.ParseGlob(pattern)
	if err != nil {
		return nil, fmt.Errorf("parse glob: %w", err)
	}
	for _, t := range tmpl.Templates() {
		slog.Info("registered component set template file", "file", t.Name())
	}

	return tmpl, nil

}
