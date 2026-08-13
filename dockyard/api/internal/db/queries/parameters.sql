-- name: CreateParameter :one
INSERT INTO parameters (image_version_id, env_var_name, type, default_value, required, description)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: ListParametersByVersionID :many
SELECT * FROM parameters
WHERE image_version_id = $1
ORDER BY env_var_name;