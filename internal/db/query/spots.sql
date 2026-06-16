-- name: ListSpots :many
SELECT * FROM spots ORDER BY id;

-- name: GetSpot :one
SELECT * FROM spots WHERE id = $1;

-- name: UpsertSpot :exec
INSERT INTO spots (
    id, name, region, latitude, longitude, orientation_degrees, buoy_ids, timezone,
    ideal_swell_dir_min, ideal_swell_dir_max, ideal_period_min, preferred_wind_dir, tide_preference
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
)
ON CONFLICT (id) DO UPDATE SET
    name                = EXCLUDED.name,
    region              = EXCLUDED.region,
    latitude            = EXCLUDED.latitude,
    longitude           = EXCLUDED.longitude,
    orientation_degrees = EXCLUDED.orientation_degrees,
    buoy_ids            = EXCLUDED.buoy_ids,
    timezone            = EXCLUDED.timezone,
    ideal_swell_dir_min = EXCLUDED.ideal_swell_dir_min,
    ideal_swell_dir_max = EXCLUDED.ideal_swell_dir_max,
    ideal_period_min    = EXCLUDED.ideal_period_min,
    preferred_wind_dir  = EXCLUDED.preferred_wind_dir,
    tide_preference     = EXCLUDED.tide_preference,
    updated_at          = now();
