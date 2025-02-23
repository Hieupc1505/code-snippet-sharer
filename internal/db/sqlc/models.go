// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"time"
)

type CodeSnippet struct {
	ID          int32     `json:"id"`
	Slug        string    `json:"slug"`
	Title       string    `json:"title"`
	Snippet     string    `json:"snippet"`
	Lang        string    `json:"lang"`
	Public      bool      `json:"public"`
	ViewCount   int32     `json:"view_count"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
}
