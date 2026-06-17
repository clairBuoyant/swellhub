// Package store is the persistence boundary the rest of the app depends on. It
// wraps the sqlc-generated queries and maps between domain types and DB rows so
// callers never touch pgtype or sqlc directly.
package store

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/clairBuoyant/swellhub/internal/db/sqlc"
	"github.com/clairBuoyant/swellhub/internal/spot"
	"github.com/clairBuoyant/swellhub/pkg/noaa"
)

// ErrNotFound is returned when a requested row does not exist.
var ErrNotFound = errors.New("not found")

// Observation is a stored NDBC observation with nullable readings preserved as
// pointers (a nil reading is "missing", distinct from a real zero).
type Observation struct {
	StationID           string
	ObservedAt          time.Time
	WaveHeightM         *float64
	DominantWavePeriodS *float64
	AverageWavePeriodS  *float64
	WaveDirectionDeg    *int32
	WindDirectionDeg    *int32
	WindSpeedMps        *float64
	WindGustMps         *float64
	AirTemperatureC     *float64
	WaterTemperatureC   *float64
	SeaLevelPressureHpa *float64
}

// Feedback is a surfer's rating plus the full conditions snapshot that produced
// our read. Conditions is opaque JSON so the exact shape can evolve.
type Feedback struct {
	SpotID             string
	ObservedAt         time.Time
	Conditions         []byte
	ComputedRating     *string
	ComputedConfidence *float64
	SpotConfigVersion  *string
	UserRating         string
	Note               *string
}

// Store is the set of persistence operations the app needs.
type Store interface {
	ListSpots(ctx context.Context) ([]spot.Spot, error)
	GetSpot(ctx context.Context, id string) (spot.Spot, error)
	UpsertSpot(ctx context.Context, s spot.Spot) error
	UpsertObservation(ctx context.Context, stationID string, o noaa.MeteorologicalObservation) (int64, error)
	LatestObservation(ctx context.Context, stationID string) (Observation, error)
	InsertFeedback(ctx context.Context, f Feedback) (int64, error)
}

// PgStore is the Postgres-backed Store.
type PgStore struct {
	q *sqlc.Queries
}

var _ Store = (*PgStore)(nil)

// New returns a PgStore backed by the given pool.
func New(pool *pgxpool.Pool) *PgStore {
	return &PgStore{q: sqlc.New(pool)}
}

func (s *PgStore) ListSpots(ctx context.Context) ([]spot.Spot, error) {
	rows, err := s.q.ListSpots(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]spot.Spot, 0, len(rows))
	for _, r := range rows {
		out = append(out, spotFromRow(r))
	}
	return out, nil
}

func (s *PgStore) GetSpot(ctx context.Context, id string) (spot.Spot, error) {
	r, err := s.q.GetSpot(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return spot.Spot{}, ErrNotFound
	}
	if err != nil {
		return spot.Spot{}, err
	}
	return spotFromRow(r), nil
}

func (s *PgStore) UpsertSpot(ctx context.Context, sp spot.Spot) error {
	return s.q.UpsertSpot(ctx, sqlc.UpsertSpotParams{
		ID:                 sp.ID,
		Name:               sp.Name,
		Region:             sp.Region,
		Latitude:           sp.Latitude,
		Longitude:          sp.Longitude,
		OrientationDegrees: int32(sp.OrientationDegrees),
		BuoyIds:            sp.BuoyIDs,
		Timezone:           sp.Timezone,
		IdealSwellDirMin:   int32(sp.IdealSwellDirMin),
		IdealSwellDirMax:   int32(sp.IdealSwellDirMax),
		IdealPeriodMin:     sp.IdealPeriodMin,
		PreferredWindDir:   int32(sp.PreferredWindDir),
		TidePreference:     string(sp.TidePreference),
	})
}

// UpsertObservation inserts an observation, returning the number of rows
// inserted (0 if it already existed).
func (s *PgStore) UpsertObservation(ctx context.Context, stationID string, o noaa.MeteorologicalObservation) (int64, error) {
	// TODO(#220): noaa.parseValue maps the NDBC "MM" sentinel to 0, so every
	// reading is stored as present. Make the noaa layer missing-aware so these
	// nullable columns reflect true absence.
	return s.q.UpsertObservation(ctx, sqlc.UpsertObservationParams{
		StationID:           stationID,
		ObservedAt:          ts(o.Datetime),
		WaveHeightM:         f8(o.WaveHeight),
		DominantWavePeriodS: f8(o.DominantWavePeriod),
		AverageWavePeriodS:  f8(o.AverageWavePeriod),
		WaveDirectionDeg:    i4(o.WaveDirection),
		WindDirectionDeg:    i4(o.WindDirection),
		WindSpeedMps:        f8(o.WindSpeed),
		WindGustMps:         f8(o.WindGust),
		AirTemperatureC:     f8(o.AirTemperature),
		WaterTemperatureC:   f8(o.WaterTemperature),
		SeaLevelPressureHpa: f8(o.SeaLevelPressure),
	})
}

func (s *PgStore) LatestObservation(ctx context.Context, stationID string) (Observation, error) {
	r, err := s.q.LatestObservation(ctx, stationID)
	if errors.Is(err, pgx.ErrNoRows) {
		return Observation{}, ErrNotFound
	}
	if err != nil {
		return Observation{}, err
	}
	return Observation{
		StationID:           r.StationID,
		ObservedAt:          r.ObservedAt.Time,
		WaveHeightM:         f8ptr(r.WaveHeightM),
		DominantWavePeriodS: f8ptr(r.DominantWavePeriodS),
		AverageWavePeriodS:  f8ptr(r.AverageWavePeriodS),
		WaveDirectionDeg:    i4ptr(r.WaveDirectionDeg),
		WindDirectionDeg:    i4ptr(r.WindDirectionDeg),
		WindSpeedMps:        f8ptr(r.WindSpeedMps),
		WindGustMps:         f8ptr(r.WindGustMps),
		AirTemperatureC:     f8ptr(r.AirTemperatureC),
		WaterTemperatureC:   f8ptr(r.WaterTemperatureC),
		SeaLevelPressureHpa: f8ptr(r.SeaLevelPressureHpa),
	}, nil
}

func (s *PgStore) InsertFeedback(ctx context.Context, f Feedback) (int64, error) {
	row, err := s.q.InsertFeedback(ctx, sqlc.InsertFeedbackParams{
		SpotID:             f.SpotID,
		ObservedAt:         ts(f.ObservedAt),
		Conditions:         f.Conditions,
		ComputedRating:     txt(f.ComputedRating),
		ComputedConfidence: f8p(f.ComputedConfidence),
		SpotConfigVersion:  txt(f.SpotConfigVersion),
		UserRating:         f.UserRating,
		Note:               txt(f.Note),
	})
	if err != nil {
		return 0, err
	}
	return row.ID, nil
}

func spotFromRow(r sqlc.Spot) spot.Spot {
	return spot.Spot{
		ID:                 r.ID,
		Name:               r.Name,
		Region:             r.Region,
		Latitude:           r.Latitude,
		Longitude:          r.Longitude,
		OrientationDegrees: int(r.OrientationDegrees),
		BuoyIDs:            r.BuoyIds,
		Timezone:           r.Timezone,
		IdealSwellDirMin:   int(r.IdealSwellDirMin),
		IdealSwellDirMax:   int(r.IdealSwellDirMax),
		IdealPeriodMin:     r.IdealPeriodMin,
		PreferredWindDir:   int(r.PreferredWindDir),
		TidePreference:     spot.TidePreference(r.TidePreference),
	}
}

// --- pgtype conversion helpers ---

func ts(t time.Time) pgtype.Timestamptz { return pgtype.Timestamptz{Time: t, Valid: true} }

func f8(v float32) pgtype.Float8 { return pgtype.Float8{Float64: float64(v), Valid: true} }

func i4(v int16) pgtype.Int4 { return pgtype.Int4{Int32: int32(v), Valid: true} }

func f8p(p *float64) pgtype.Float8 {
	if p == nil {
		return pgtype.Float8{}
	}
	return pgtype.Float8{Float64: *p, Valid: true}
}

func txt(p *string) pgtype.Text {
	if p == nil {
		return pgtype.Text{}
	}
	return pgtype.Text{String: *p, Valid: true}
}

func f8ptr(v pgtype.Float8) *float64 {
	if !v.Valid {
		return nil
	}
	f := v.Float64
	return &f
}

func i4ptr(v pgtype.Int4) *int32 {
	if !v.Valid {
		return nil
	}
	n := v.Int32
	return &n
}
