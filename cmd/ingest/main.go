// Command ingest fetches realtime NDBC observations for the buoys mapped to the
// configured spots and upserts them into Postgres. Intended to run on a
// schedule (see .github/workflows/ingest.yml).
//
// Usage: DB_URL=postgres://... go run ./cmd/ingest
package main

import (
	"context"
	"os"

	"github.com/clairBuoyant/swellhub/internal/db"
	"github.com/clairBuoyant/swellhub/internal/ingest"
	"github.com/clairBuoyant/swellhub/internal/spot"
	"github.com/clairBuoyant/swellhub/internal/store"
	"github.com/clairBuoyant/swellhub/pkg/log"
	"github.com/clairBuoyant/swellhub/pkg/noaa"
)

func main() {
	logger := log.InitLoggerJSON()

	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		logger.Error("DB_URL is required")
		os.Exit(1)
	}

	ctx := context.Background()
	pool, err := db.NewPool(ctx, dsn)
	if err != nil {
		logger.Error("connect postgres", "error", err)
		os.Exit(1)
	}
	defer pool.Close()

	stations := mappedStations()
	n, err := ingest.Run(ctx, store.New(pool), noaa.Realtime, stations, logger)
	if err != nil {
		logger.Error("ingest completed with failures", "stations", len(stations), "observations_inserted", n, "error", err)
		pool.Close()
		os.Exit(1)
	}
	logger.Info("ingest complete", "stations", len(stations), "observations_inserted", n)
}

// mappedStations returns the de-duplicated NDBC station IDs across all spots,
// preserving first-seen order.
func mappedStations() []string {
	seen := make(map[string]bool)
	var ids []string
	for _, s := range spot.All() {
		for _, id := range s.BuoyIDs {
			if !seen[id] {
				seen[id] = true
				ids = append(ids, id)
			}
		}
	}
	return ids
}
