// Package spot defines the core surf-break domain entity and a small seed of
// hand-tuned NY/NJ spots. A Spot carries the static config used to interpret
// offshore buoy conditions for that break (orientation, mapped buoys, ideal
// windows). Conditions themselves are fetched separately (see pkg/noaa).
package spot

// TidePreference describes the tide stage a break tends to favor. Used by the
// rating engine (M1-1); kept coarse for now.
type TidePreference string

const (
	TideAny  TidePreference = "any"
	TideLow  TidePreference = "low"
	TideMid  TidePreference = "mid"
	TideHigh TidePreference = "high"
)

// Spot is a surf break and the static config used to interpret conditions there.
type Spot struct {
	ID                 string // stable slug, e.g. "rockaway-90th"
	Name               string
	Region             string // "NY", "NJ"
	Latitude           float64
	Longitude          float64
	OrientationDegrees int      // azimuth the beach faces (seaward normal)
	BuoyIDs            []string // prioritized NDBC stations; first reachable wins
	Timezone           string   // IANA tz, e.g. "America/New_York"

	// Rating config windows (consumed in M1-1). Hand-tuned and intentionally
	// simple; OSM-derived orientation and learned windows come later.
	IdealSwellDirMin int            // degrees
	IdealSwellDirMax int            // degrees
	IdealPeriodMin   float64        // seconds
	PreferredWindDir int            // offshore source direction (deg wind comes from)
	TidePreference   TidePreference // preferred tide stage
}

// seed is the hand-tuned starter set. Buoy mappings and orientation are
// approximate and meant to be validated against real sessions via the
// feedback loop.
var seed = []Spot{
	{
		ID:                 "rockaway-90th",
		Name:               "Rockaway Beach 90th St",
		Region:             "NY",
		Latitude:           40.5828,
		Longitude:          -73.8186,
		OrientationDegrees: 180, // faces ~south
		BuoyIDs:            []string{"44065", "44025"},
		Timezone:           "America/New_York",
		IdealSwellDirMin:   150,
		IdealSwellDirMax:   210,
		IdealPeriodMin:     8,
		PreferredWindDir:   0, // offshore from the north (orientation ± 180)
		TidePreference:     TideAny,
	},
	{
		ID:                 "manasquan",
		Name:               "Manasquan Inlet",
		Region:             "NJ",
		Latitude:           40.1037,
		Longitude:          -74.0354,
		OrientationDegrees: 100, // faces ~east-southeast
		BuoyIDs:            []string{"44091", "44065", "44025"},
		Timezone:           "America/New_York",
		IdealSwellDirMin:   70,
		IdealSwellDirMax:   140,
		IdealPeriodMin:     8,
		PreferredWindDir:   280, // offshore from the west
		TidePreference:     TideMid,
	},
}

// All returns a copy of the configured spots.
func All() []Spot {
	out := make([]Spot, len(seed))
	copy(out, seed)
	return out
}

// ByID returns the spot with the given slug.
func ByID(id string) (Spot, bool) {
	for _, s := range seed {
		if s.ID == id {
			return s, true
		}
	}
	return Spot{}, false
}
