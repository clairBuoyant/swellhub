-- name: UpsertObservation :exec
INSERT INTO observations (
    station_id, observed_at, wave_height_m, dominant_wave_period_s, average_wave_period_s,
    wave_direction_deg, wind_direction_deg, wind_speed_mps, wind_gust_mps,
    air_temperature_c, water_temperature_c, sea_level_pressure_hpa
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
)
ON CONFLICT (station_id, observed_at) DO NOTHING;

-- name: LatestObservation :one
SELECT * FROM observations
WHERE station_id = $1
ORDER BY observed_at DESC
LIMIT 1;
