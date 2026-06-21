//go:build integration

package store_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/clairBuoyant/swellhub/internal/db"
	"github.com/clairBuoyant/swellhub/internal/spot"
	"github.com/clairBuoyant/swellhub/internal/store"
	"github.com/clairBuoyant/swellhub/pkg/noaa"
)

// testStore connects to TEST_DB_URL, applies migrations, and returns a Store. These
// tests are gated by the `integration` build tag, so a plain `go test ./...`
// never runs them — only `task db:test` / CI (with -tags=integration) do, which
// point TEST_DB_URL at a throwaway database. Skipped if it is unset.
func testStore(t *testing.T) (*store.PgStore, *pgxpool.Pool) {
	t.Helper()
	dsn := os.Getenv("TEST_DB_URL")
	if dsn == "" {
		t.Skip("TEST_DB_URL not set; skipping store integration tests")
	}
	if err := db.Migrate(dsn); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	pool, err := db.NewPool(context.Background(), dsn)
	if err != nil {
		t.Fatalf("pool: %v", err)
	}
	t.Cleanup(pool.Close)
	return store.New(pool), pool
}

func TestSpotsRoundTrip(t *testing.T) {
	st, _ := testStore(t)
	ctx := context.Background()

	want := spot.Spot{
		ID: "test-spot", Name: "Test", Region: "NY",
		Latitude: 40, Longitude: -73, OrientationDegrees: 180,
		BuoyIDs: []string{"44065", "44025"}, Timezone: "America/New_York",
		IdealSwellDirMin: 150, IdealSwellDirMax: 210, IdealPeriodMin: 8,
		PreferredWindDir: 0, TidePreference: spot.TideAny,
	}
	if err := st.UpsertSpot(ctx, want); err != nil {
		t.Fatalf("UpsertSpot: %v", err)
	}

	got, err := st.GetSpot(ctx, "test-spot")
	if err != nil {
		t.Fatalf("GetSpot: %v", err)
	}
	if got.Name != want.Name || got.OrientationDegrees != want.OrientationDegrees ||
		len(got.BuoyIDs) != 2 || got.TidePreference != spot.TideAny {
		t.Errorf("round-trip mismatch: got %+v", got)
	}

	if _, err := st.GetSpot(ctx, "missing"); err != store.ErrNotFound {
		t.Errorf("GetSpot(missing) err = %v, want ErrNotFound", err)
	}

	spots, err := st.ListSpots(ctx)
	if err != nil {
		t.Fatalf("ListSpots: %v", err)
	}
	if len(spots) == 0 {
		t.Error("ListSpots returned none")
	}
}

func TestObservationsLatest(t *testing.T) {
	st, _ := testStore(t)
	ctx := context.Background()

	older := noaa.MeteorologicalObservation{Datetime: time.Date(2026, 6, 14, 10, 0, 0, 0, time.UTC), WaveHeight: float32Ptr(1.0)}
	newer := noaa.MeteorologicalObservation{Datetime: time.Date(2026, 6, 14, 11, 0, 0, 0, time.UTC), WaveHeight: float32Ptr(1.4)}
	for _, o := range []noaa.MeteorologicalObservation{older, newer} {
		if _, err := st.UpsertObservation(ctx, "TESTBUOY", o); err != nil {
			t.Fatalf("UpsertObservation: %v", err)
		}
	}
	// Re-upsert is idempotent (ON CONFLICT DO NOTHING) -> 0 rows inserted.
	if n, err := st.UpsertObservation(ctx, "TESTBUOY", newer); err != nil {
		t.Fatalf("re-UpsertObservation: %v", err)
	} else if n != 0 {
		t.Errorf("re-upsert inserted %d rows, want 0", n)
	}

	got, err := st.LatestObservation(ctx, "TESTBUOY")
	if err != nil {
		t.Fatalf("LatestObservation: %v", err)
	}
	if !got.ObservedAt.Equal(newer.Datetime) {
		t.Errorf("latest observed_at = %v, want %v", got.ObservedAt, newer.Datetime)
	}
	if got.WaveHeightM == nil {
		t.Fatal("latest wave height is nil")
	}
	// Stored as float32 then widened, so compare via float32 to avoid precision noise.
	if float32(*got.WaveHeightM) != 1.4 {
		t.Errorf("latest wave height = %v, want 1.4", *got.WaveHeightM)
	}
	if got.WindSpeedMps != nil {
		t.Errorf("latest missing wind speed = %v, want nil", *got.WindSpeedMps)
	}

	if _, err := st.LatestObservation(ctx, "NOBUOY"); err != store.ErrNotFound {
		t.Errorf("LatestObservation(NOBUOY) err = %v, want ErrNotFound", err)
	}
}

func float32Ptr(v float32) *float32 { return &v }

func TestInsertFeedback(t *testing.T) {
	st, _ := testStore(t)
	ctx := context.Background()

	// Needs a spot to satisfy the FK.
	if err := st.UpsertSpot(ctx, spot.Spot{
		ID: "fb-spot", Name: "FB", Region: "NJ", Timezone: "America/New_York",
		BuoyIDs: []string{"44091"}, TidePreference: spot.TideMid,
	}); err != nil {
		t.Fatalf("UpsertSpot: %v", err)
	}

	rating := "good"
	id, err := st.InsertFeedback(ctx, store.Feedback{
		SpotID:         "fb-spot",
		ObservedAt:     time.Date(2026, 6, 14, 12, 0, 0, 0, time.UTC),
		Conditions:     []byte(`{"waveHeightM":1.2}`),
		ComputedRating: &rating,
		UserRating:     "fair",
	})
	if err != nil {
		t.Fatalf("InsertFeedback: %v", err)
	}
	if id == 0 {
		t.Error("expected non-zero feedback id")
	}
}
