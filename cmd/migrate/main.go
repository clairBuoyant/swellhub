// Command migrate runs database migrations and, on "up", seeds the spot config
// from the in-repo source of truth (internal/spot). Idempotent; safe to re-run.
//
// Usage: DATABASE_URL=postgres://... go run ./cmd/migrate [up|down|status|reset]
package main

import (
	"context"
	"log"
	"os"

	"github.com/clairBuoyant/swellhub/internal/db"
	"github.com/clairBuoyant/swellhub/internal/spot"
	"github.com/clairBuoyant/swellhub/internal/store"
)

func main() {
	command := "up"
	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is required")
	}

	if err := db.RunGoose(dsn, command); err != nil {
		log.Fatal(err)
	}

	// Seeding only makes sense after applying migrations.
	if command != "up" {
		return
	}

	ctx := context.Background()
	pool, err := db.NewPool(ctx, dsn)
	if err != nil {
		log.Fatalf("connect: %v", err)
	}
	defer pool.Close()

	st := store.New(pool)
	spots := spot.All()
	for _, s := range spots {
		if err := st.UpsertSpot(ctx, s); err != nil {
			log.Fatalf("seed spot %q: %v", s.ID, err)
		}
	}
	log.Printf("seeded %d spots", len(spots))
}
