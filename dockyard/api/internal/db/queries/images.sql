-- name: ListImages :many
SELECT id, name, registry, description, source_url, logo_url, created_at, updated_at
FROM images
ORDER BY name
LIMIT $1 OFFSET $2;

-- name: CreateImage :one
INSERT INTO images (name, registry, description, source_url, logo_url)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, name, registry, description, source_url, logo_url, created_at, updated_at;

-- name: GetImageByID :one
SELECT id, name, registry, description, source_url, logo_url, created_at, updated_at
FROM images
WHERE id = $1;