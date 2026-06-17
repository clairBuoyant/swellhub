// Package ingest fetches realtime NDBC observations for a set of stations and
// upserts them, starting the durable observation history (the "data clock").
package ingest

import (
	"context"
	"log/slog"

	"github.com/clairBuoyant/swellhub/pkg/noaa"
)

// Upserter persists a single observation, returning rows inserted (0 if it
// already existed). *store.PgStore satisfies it.
type Upserter interface {
	UpsertObservation(ctx context.Context, stationID string, o noaa.MeteorologicalObservation) (int64, error)
}

// RealtimeFunc fetches realtime observations for a station (noaa.Realtime).
type RealtimeFunc func(stationID string, dataset noaa.RealtimeDataset) ([]noaa.MeteorologicalObservation, error)

// Run fetches realtime observations for each station and upserts every returned
// row (idempotent on station+datetime). It is best-effort: an error for one
// station is logged and does not abort the others. Returns the number of
// observations newly inserted (dupes count as 0).
func Run(ctx context.Context, store Upserter, fetch RealtimeFunc, stationIDs []string, logger *slog.Logger) int {
	var inserted int
	for _, id := range stationIDs {
		obs, err := fetch(id, noaa.TXT)
		if err != nil {
			logger.Warn("buoy fetch failed", "buoy", id, "error", err)
			continue
		}
		for _, o := range obs {
			n, err := store.UpsertObservation(ctx, id, o)
			if err != nil {
				logger.Warn("upsert failed", "buoy", id, "observed_at", o.Datetime, "error", err)
				continue
			}
			inserted += int(n)
		}
	}
	return inserted
}
