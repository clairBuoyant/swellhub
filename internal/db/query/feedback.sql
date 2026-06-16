-- name: InsertFeedback :one
INSERT INTO feedback (
    spot_id, observed_at, conditions, computed_rating, computed_confidence,
    spot_config_version, user_rating, note
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;
