-- name: Add :exec
INSERT INTO code_snippets (slug, title, snippet, lang, public, view_count)
VALUES ($1, $2, $3, $4, $5, $6);


-- name: GetBySlug :one
SELECT * FROM code_snippets WHERE slug = $1 LIMIT 1;


-- name: UpdateViewCount :exec
UPDATE code_snippets
SET view_count = view_count + 1
WHERE slug = $1;

-- name: GetPublicSnippets :many
SELECT * FROM code_snippets
WHERE public = TRUE;


-- name: GetRecentPosts :many
SELECT * FROM code_snippets
ORDER BY created_time DESC
LIMIT 4;