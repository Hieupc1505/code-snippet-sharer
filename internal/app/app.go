package app

import (
	"errors"
	"github.com/jackc/pgx/v5"
)

const (
	Name  = "GopherHunt"
	Title = "GopherHunt - Pairing talented Go Engineers with great companies"
)

var (
	ErrInvalidInput        = errors.New("invalid input")
	ErrUnauthorised        = errors.New("unauthorised")
	ErrForbidden           = errors.New("forbidden")
	ErrUnsupportedLanguage = errors.New("unsupported language")
	ErrEmailInvalidLength  = errors.New("invalid title length")
	ErrNotFound            = pgx.ErrNoRows
)
