package account

import (
	"context"
	"errors"
	"fmt"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
	db "s-coder-snippet-sharder/internal/db/sqlc"
	"s-coder-snippet-sharder/pkg/config"
)

var (
	ErrEmailInvalidLen    = errors.New("invalid email length")
	ErrPasswordInvalidLen = errors.New("invalid password length")
)

type Read interface {
}

type Writer interface {
}

type ReadWriter interface {
	Read
	Writer
}

type Service struct {
	repo ReadWriter
	Db   db.DBTX
}

func (s *Service) RegisterAuthService() {
	gothic.Store = nil

	goth.UseProviders(
		github.New(
			config.Envs.GithubClientID,
			config.Envs.GithubClientSecret,
			buildCallbackURL("github"),
		),
	)
}

func buildCallbackURL(provider string) string {
	return fmt.Sprintf("%s:%s/api/auth/%s/callback", config.Envs.PublicHost, config.Envs.Port, provider)
}

func NewService(ctx context.Context, repo ReadWriter) (*Service, error) {
	return &Service{repo: repo}, nil
}
