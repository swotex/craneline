-- name: CreateImageVersion :one
INSERT INTO image_versions (image_id, tag, digest, is_latest, released_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: ListVersionsByImageID :many
SELECT * FROM image_versions
WHERE image_id = $1
ORDER BY created_at DESC;