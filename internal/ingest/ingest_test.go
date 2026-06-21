package ingest

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"testing"
	"time"

	"github.com/clairBuoyant/swellhub/pkg/noaa"
)

type fakeStore struct {
	upserts int
	failOn  string
}

func (f *fakeStore) UpsertObservation(_ context.Context, stationID string, _ noaa.MeteorologicalObservation) (int64, error) {
	if stationID == f.failOn {
		return 0, errors.New("db error")
	}
	f.upserts++
	return 1, nil
}

func discardLogger() *slog.Logger { return slog.New(slog.NewTextHandler(io.Discard, nil)) }

func TestRun(t *testing.T) {
	two := []noaa.MeteorologicalObservation{
		{Datetime: time.Now()},
		{Datetime: time.Now().Add(-time.Hour)},
	}
	fetch := func(id string, _ noaa.RealtimeDataset) ([]noaa.MeteorologicalObservation, error) {
		switch id {
		case "GOOD":
			return two, nil
		case "EMPTY":
			return nil, nil
		}
		return nil, nil
	}

	st := &fakeStore{}
	// GOOD -> 2 upserts, EMPTY -> 0.
	n, err := Run(context.Background(), st, fetch, []string{"GOOD", "EMPTY"}, discardLogger())
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	if n != 2 {
		t.Errorf("upserted = %d, want 2", n)
	}
	if st.upserts != 2 {
		t.Errorf("store.upserts = %d, want 2", st.upserts)
	}
}

func TestRunContinuesPastUpsertError(t *testing.T) {
	obs := []noaa.MeteorologicalObservation{{Datetime: time.Now()}}
	fetch := func(string, noaa.RealtimeDataset) ([]noaa.MeteorologicalObservation, error) {
		return obs, nil
	}
	st := &fakeStore{failOn: "BAD"}
	// BAD upsert fails (logged, skipped); OK succeeds.
	n, err := Run(context.Background(), st, fetch, []string{"BAD", "OK"}, discardLogger())
	if n != 1 {
		t.Errorf("upserted = %d, want 1", n)
	}
	if err == nil {
		t.Fatal("Run error = nil, want partial-failure error")
	}
}

func TestRunReportsFetchErrorAfterContinuing(t *testing.T) {
	fetchErr := errors.New("ndbc down")
	fetch := func(id string, _ noaa.RealtimeDataset) ([]noaa.MeteorologicalObservation, error) {
		if id == "BAD" {
			return nil, fetchErr
		}
		return []noaa.MeteorologicalObservation{{Datetime: time.Now()}}, nil
	}

	n, err := Run(context.Background(), &fakeStore{}, fetch, []string{"BAD", "OK"}, discardLogger())
	if n != 1 {
		t.Errorf("inserted = %d, want 1", n)
	}
	if !errors.Is(err, fetchErr) {
		t.Fatalf("Run error = %v, want wrapped fetch error", err)
	}
}
