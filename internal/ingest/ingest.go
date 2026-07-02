// Package ingest fetches realtime NDBC observations for a set of stations and
// upserts them, starting the durable observation history (the "data clock").
package ingest

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/clairBuoyant/swellhub/pkg/noaa"
)

// Upserter persists a single observation, returning rows inserted (0 if it
// already existed). *store.PgStore satisfies it.
type Upserter interface {
	UpsertObservation(ctx context.Context, stationID string, o noaa.MeteorologicalObservation) (int64, error)
}

// RealtimeFunc fetches realtime observations for a station (noaa.Realtime).
type RealtimeFunc func(ctx context.Context, stationID string, dataset noaa.RealtimeDataset) ([]noaa.MeteorologicalObservation, error)

// Run fetches realtime observations for each station and upserts every returned
// row (idempotent on station+datetime). It is best-effort: an error for one
// station is logged and does not abort the others. Returns the number of
// observations newly inserted (dupes count as 0). All stations are attempted;
// any fetch or persistence failures are joined and returned after processing.
func Run(ctx context.Context, store Upserter, fetch RealtimeFunc, stationIDs []string, logger *slog.Logger) (int, error) {
	var inserted int
	var failures []error
	for _, id := range stationIDs {
		obs, err := fetch(ctx, id, noaa.TXT)
		if err != nil {
			logger.Warn("buoy fetch failed", "buoy", id, "error", err)
			failures = append(failures, fmt.Errorf("fetch buoy %s: %w", id, err))
			continue
		}
		for _, o := range obs {
			n, err := store.UpsertObservation(ctx, id, o)
			if err != nil {
				logger.Warn("upsert failed", "buoy", id, "observed_at", o.Datetime, "error", err)
				failures = append(failures, fmt.Errorf("upsert buoy %s at %s: %w", id, o.Datetime, err))
				continue
			}
			inserted += int(n)
		}
	}
	return inserted, errors.Join(failures...)
}
