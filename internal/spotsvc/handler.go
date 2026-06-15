// Package spotsvc implements the Connect SpotService: it serves configured
// spots and joins each with the latest conditions from its mapped buoys.
package spotsvc

import (
	"context"
	"fmt"
	"log/slog"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/types/known/timestamppb"

	spotv1 "github.com/clairBuoyant/swellhub/gen/clairbuoyant/spot/v1"
	"github.com/clairBuoyant/swellhub/pkg/noaa"
	"github.com/clairBuoyant/swellhub/pkg/spot"
)

// realtimeFunc fetches realtime observations for a station. Injected so tests
// can stub NDBC.
type realtimeFunc func(stationID string, dataset noaa.RealtimeDataset) ([]noaa.MeteorologicalObservation, error)

// Service implements spotv1connect.SpotServiceHandler.
type Service struct {
	logger   *slog.Logger
	realtime realtimeFunc
}

// New returns a Service backed by live NDBC data.
func New(logger *slog.Logger) *Service {
	return &Service{logger: logger, realtime: noaa.Realtime}
}

// ListSpots returns all configured spots (config only).
func (s *Service) ListSpots(
	_ context.Context,
	_ *connect.Request[spotv1.ListSpotsRequest],
) (*connect.Response[spotv1.ListSpotsResponse], error) {
	spots := spot.All()
	out := make([]*spotv1.Spot, 0, len(spots))
	for _, sp := range spots {
		out = append(out, toProtoSpot(sp))
	}
	return connect.NewResponse(&spotv1.ListSpotsResponse{Spots: out}), nil
}

// GetSpot returns a spot joined with the latest conditions from its primary
// reachable buoy. Unknown id yields CodeNotFound.
func (s *Service) GetSpot(
	_ context.Context,
	req *connect.Request[spotv1.GetSpotRequest],
) (*connect.Response[spotv1.GetSpotResponse], error) {
	sp, ok := spot.ByID(req.Msg.GetId())
	if !ok {
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("spot %q not found", req.Msg.GetId()))
	}

	return connect.NewResponse(&spotv1.GetSpotResponse{
		SpotConditions: &spotv1.SpotConditions{
			Spot:       toProtoSpot(sp),
			Conditions: s.latestConditions(sp),
		},
	}), nil
}

// latestConditions returns the most recent observation from the first mapped
// buoy that responds with data, or nil if none do.
func (s *Service) latestConditions(sp spot.Spot) *spotv1.Conditions {
	for _, buoyID := range sp.BuoyIDs {
		obs, err := s.realtime(buoyID, noaa.TXT)
		if err != nil {
			s.logger.Warn("buoy fetch failed", "spot", sp.ID, "buoy", buoyID, "error", err)
			continue
		}
		if len(obs) == 0 {
			continue
		}
		// NDBC realtime files are newest-first, so index 0 is the latest.
		return toProtoConditions(buoyID, obs[0])
	}
	return nil
}

func toProtoSpot(sp spot.Spot) *spotv1.Spot {
	return &spotv1.Spot{
		Id:                 sp.ID,
		Name:               sp.Name,
		Region:             sp.Region,
		Latitude:           sp.Latitude,
		Longitude:          sp.Longitude,
		OrientationDegrees: int32(sp.OrientationDegrees),
		BuoyIds:            sp.BuoyIDs,
		Timezone:           sp.Timezone,
	}
}

func toProtoConditions(buoyID string, o noaa.MeteorologicalObservation) *spotv1.Conditions {
	// TODO(@kylejb): noaa.parseValue maps the NDBC "MM" sentinel to 0, so we
	// can't yet distinguish missing from a real zero. Until that layer is made
	// missing-aware, every field is reported as present.
	return &spotv1.Conditions{
		BuoyId:              buoyID,
		ObservedAt:          timestamppb.New(o.Datetime),
		WaveHeightM:         f64p(o.WaveHeight),
		DominantWavePeriodS: f64p(o.DominantWavePeriod),
		AverageWavePeriodS:  f64p(o.AverageWavePeriod),
		WaveDirectionDeg:    i32p(o.WaveDirection),
		WindDirectionDeg:    i32p(o.WindDirection),
		WindSpeedMps:        f64p(o.WindSpeed),
		WindGustMps:         f64p(o.WindGust),
		AirTemperatureC:     f64p(o.AirTemperature),
		WaterTemperatureC:   f64p(o.WaterTemperature),
		SeaLevelPressureHpa: f64p(o.SeaLevelPressure),
	}
}

func f64p(v float32) *float64 { f := float64(v); return &f }
func i32p(v int16) *int32     { i := int32(v); return &i }
