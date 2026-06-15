package http

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"testing"
	"time"

	"connectrpc.com/connect"

	spotv1 "github.com/clairBuoyant/swellhub/gen/clairbuoyant/spot/v1"
	"github.com/clairBuoyant/swellhub/pkg/noaa"
)

func testService(fn realtimeFunc) *spotService {
	return &spotService{
		logger:   slog.New(slog.NewTextHandler(io.Discard, nil)),
		realtime: fn,
	}
}

func TestListSpots(t *testing.T) {
	svc := testService(func(string, noaa.RealtimeDataset) ([]noaa.MeteorologicalObservation, error) {
		return nil, nil
	})

	resp, err := svc.ListSpots(context.Background(), connect.NewRequest(&spotv1.ListSpotsRequest{}))
	if err != nil {
		t.Fatalf("ListSpots: %v", err)
	}
	if len(resp.Msg.GetSpots()) == 0 {
		t.Fatal("expected seeded spots, got none")
	}

	var found bool
	for _, s := range resp.Msg.GetSpots() {
		if s.GetId() == "rockaway-90th" {
			found = true
		}
	}
	if !found {
		t.Error("expected rockaway-90th in ListSpots")
	}
}

func TestGetSpotNotFound(t *testing.T) {
	svc := testService(func(string, noaa.RealtimeDataset) ([]noaa.MeteorologicalObservation, error) {
		return nil, nil
	})

	_, err := svc.GetSpot(context.Background(), connect.NewRequest(&spotv1.GetSpotRequest{Id: "does-not-exist"}))
	if err == nil {
		t.Fatal("expected error for unknown spot")
	}
	if got := connect.CodeOf(err); got != connect.CodeNotFound {
		t.Errorf("error code = %v, want %v", got, connect.CodeNotFound)
	}
}

func TestGetSpotConditions(t *testing.T) {
	obs := noaa.MeteorologicalObservation{
		Datetime:   time.Date(2026, 6, 14, 12, 0, 0, 0, time.UTC),
		WaveHeight: 1.5,
		WindSpeed:  4.0,
	}

	// rockaway-90th maps buoys ["44065", "44025"].
	tests := []struct {
		name     string
		fn       realtimeFunc
		wantBuoy string
		wantNil  bool
	}{
		{
			name: "primary buoy returns data",
			fn: func(id string, _ noaa.RealtimeDataset) ([]noaa.MeteorologicalObservation, error) {
				if id == "44065" {
					return []noaa.MeteorologicalObservation{obs}, nil
				}
				return nil, nil
			},
			wantBuoy: "44065",
		},
		{
			name: "falls back to next buoy when primary errors",
			fn: func(id string, _ noaa.RealtimeDataset) ([]noaa.MeteorologicalObservation, error) {
				switch id {
				case "44065":
					return nil, errors.New("ndbc unavailable")
				case "44025":
					return []noaa.MeteorologicalObservation{obs}, nil
				}
				return nil, nil
			},
			wantBuoy: "44025",
		},
		{
			name: "no usable data from any buoy",
			fn: func(string, noaa.RealtimeDataset) ([]noaa.MeteorologicalObservation, error) {
				return nil, nil
			},
			wantNil: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc := testService(tc.fn)

			resp, err := svc.GetSpot(context.Background(), connect.NewRequest(&spotv1.GetSpotRequest{Id: "rockaway-90th"}))
			if err != nil {
				t.Fatalf("GetSpot: %v", err)
			}

			conds := resp.Msg.GetSpotConditions().GetConditions()
			if tc.wantNil {
				if conds != nil {
					t.Fatalf("expected nil conditions, got %+v", conds)
				}
				return
			}
			if conds == nil {
				t.Fatal("expected conditions, got nil")
			}
			if conds.GetBuoyId() != tc.wantBuoy {
				t.Errorf("buoy = %q, want %q", conds.GetBuoyId(), tc.wantBuoy)
			}
			if conds.GetWaveHeightM() != 1.5 {
				t.Errorf("wave height = %v, want 1.5", conds.GetWaveHeightM())
			}
		})
	}
}
