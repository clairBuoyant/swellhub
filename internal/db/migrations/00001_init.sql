-- +goose Up
-- +goose StatementBegin
CREATE TABLE spots (
    id                  TEXT PRIMARY KEY,
    name                TEXT             NOT NULL,
    region              TEXT             NOT NULL,
    latitude            DOUBLE PRECISION NOT NULL,
    longitude           DOUBLE PRECISION NOT NULL,
    orientation_degrees INTEGER          NOT NULL,
    buoy_ids            TEXT[]           NOT NULL,
    timezone            TEXT             NOT NULL,
    ideal_swell_dir_min INTEGER          NOT NULL,
    ideal_swell_dir_max INTEGER          NOT NULL,
    ideal_period_min    DOUBLE PRECISION NOT NULL,
    preferred_wind_dir  INTEGER          NOT NULL,
    tide_preference     TEXT             NOT NULL,
    created_at          TIMESTAMPTZ      NOT NULL DEFAULT now(),
    updated_at          TIMESTAMPTZ      NOT NULL DEFAULT now()
);
-- +goose StatementEnd

-- +goose StatementBegin
-- Time-series of raw NDBC observations, keyed by station + observation time so
-- re-ingesting the same realtime window is idempotent.
CREATE TABLE observations (
    station_id             TEXT        NOT NULL,
    observed_at            TIMESTAMPTZ NOT NULL,
    wave_height_m          DOUBLE PRECISION,
    dominant_wave_period_s DOUBLE PRECISION,
    average_wave_period_s  DOUBLE PRECISION,
    wave_direction_deg     INTEGER,
    wind_direction_deg     INTEGER,
    wind_speed_mps         DOUBLE PRECISION,
    wind_gust_mps          DOUBLE PRECISION,
    air_temperature_c      DOUBLE PRECISION,
    water_temperature_c    DOUBLE PRECISION,
    sea_level_pressure_hpa DOUBLE PRECISION,
    created_at             TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY (station_id, observed_at)
);
-- +goose StatementEnd

-- +goose StatementBegin
-- The crown-jewel dataset: a surfer's rating paired with a full snapshot of the
-- conditions (raw inputs + our computed read) that produced it.
CREATE TABLE feedback (
    id                  BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    spot_id             TEXT        NOT NULL REFERENCES spots (id),
    observed_at         TIMESTAMPTZ NOT NULL,
    conditions          JSONB       NOT NULL,
    computed_rating     TEXT,
    computed_confidence DOUBLE PRECISION,
    spot_config_version TEXT,
    user_rating         TEXT        NOT NULL,
    note                TEXT,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT now()
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX observations_station_observed_at_idx
    ON observations (station_id, observed_at DESC);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX feedback_spot_observed_at_idx
    ON feedback (spot_id, observed_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS feedback;
DROP TABLE IF EXISTS observations;
DROP TABLE IF EXISTS spots;
-- +goose StatementEnd
