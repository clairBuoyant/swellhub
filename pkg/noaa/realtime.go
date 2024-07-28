package noaa

import (
	"encoding/json"
	"fmt"
	"time"
)

type RealtimeDataset string

// Realtime Datasets
//
//	There are nine different data sources:
//	  - data_spec     Raw Spectral Wave Data
//	  - ocean         Oceanographic Data
//	  - spec          Spectral Wave Summary Data
//	  - supl          Supplemental Measurements Data
//	  - swdir         Spectral Wave Data (alpha1)
//	  - swdir2        Spectral Wave Data (alpha2)
//	  - swr1          Spectral Wave Data (r1)
//	  - swr2          Spectral Wave Data (r2)
//	  - txt           Standard Meteorological Data
const (
	DATASPEC RealtimeDataset = "data_spec"
	OCEAN    RealtimeDataset = "ocean"
	SPEC     RealtimeDataset = "spec"
	SUPL     RealtimeDataset = "supl"
	SWDIR    RealtimeDataset = "swdir"
	SWDIR2   RealtimeDataset = "swdir2"
	SWR1     RealtimeDataset = "swr1"
	SWR2     RealtimeDataset = "swr2"
	TXT      RealtimeDataset = "txt"
)

// validDataset contains all valid dataset types.
//
// TODO(@kylejb): add support for other data sources
var validDataset = map[RealtimeDataset]bool{
	// DATASPEC: true,
	// OCEAN:    true,
	// SPEC:     true,
	// SUPL:     true,
	// SWDIR:    true,
	// SWDIR2:   true,
	// SWR1:     true,
	// SWR2:     true,
	TXT: true,
}

func (d RealtimeDataset) IsValid() bool {
	_, ok := validDataset[d]
	return ok
}

type MeteorologicalObservation struct {
	Datetime            time.Time `json:"datetime"`
	WindDirection       int16     `json:"windDirection"`
	WindSpeed           float32   `json:"windSpeed"`
	WindGust            float32   `json:"windGust"`
	WaveHeight          float32   `json:"waveHeight"`
	DominantWavePeriod  float32   `json:"dominantWavePeriod"`
	AverageWavePeriod   float32   `json:"averageWavePeriod"`
	WaveDirection       int16     `json:"waveDirection"`
	SeaLevelPressure    float32   `json:"seaLevelPressure"`
	PressureTendency    float32   `json:"pressureTendency"`
	AirTemperature      float32   `json:"airTemperature"`
	WaterTemperature    float32   `json:"waterTemperature"`
	DewpointTemperature float32   `json:"dewpointTemperature"`
	Visibility          float32   `json:"visibility"`
	Tide                float32   `json:"tide"`
}

type WaveSummaryObservation struct {
	Datetime              time.Time `json:"datetime"`
	SignificantWaveHeight float32   `json:"significantWaveHeight"`
	SwellHeight           float32   `json:"swellHeight"`
	SwellPeriod           float32   `json:"swellPeriod"`
	WindWaveHeight        float32   `json:"windWaveHeight"`
	WindWavePeriod        float32   `json:"windWavePeriod"`
	SwellDirection        string    `json:"swellDirection"`
	WindWaveDirection     float32   `json:"windWaveDirection"`
	Steepness             string    `json:"steepness"`
	AverageWavePeriod     float32   `json:"averageWavePeriod"`
	DominantWaveDirection int16     `json:"dominantWaveDirection"`
}

// Encodes time.Time object to ISODateString
func (o MeteorologicalObservation) MarshalJSON() ([]byte, error) {
	metObs := struct {
		Datetime            string  `json:"datetime"`
		WindDirection       int16   `json:"windDirection"`
		WindSpeed           float32 `json:"windSpeed"`
		WindGust            float32 `json:"windGust"`
		WaveHeight          float32 `json:"waveHeight"`
		DominantWavePeriod  float32 `json:"dominantWavePeriod"`
		AverageWavePeriod   float32 `json:"averageWavePeriod"`
		WaveDirection       int16   `json:"waveDirection"`
		SeaLevelPressure    float32 `json:"seaLevelPressure"`
		PressureTendency    float32 `json:"pressureTendency"`
		AirTemperature      float32 `json:"airTemperature"`
		WaterTemperature    float32 `json:"waterTemperature"`
		DewpointTemperature float32 `json:"dewpointTemperature"`
		Visibility          float32 `json:"visibility"`
		Tide                float32 `json:"tide"`
	}{
		Datetime:            o.Datetime.Format((time.RFC3339)),
		WindDirection:       o.WindDirection,
		WindSpeed:           o.WindSpeed,
		WindGust:            o.WindGust,
		WaveHeight:          o.WaveHeight,
		DominantWavePeriod:  o.DominantWavePeriod,
		AverageWavePeriod:   o.AverageWavePeriod,
		WaveDirection:       o.WaveDirection,
		SeaLevelPressure:    o.SeaLevelPressure,
		PressureTendency:    o.PressureTendency,
		AirTemperature:      o.AirTemperature,
		WaterTemperature:    o.WaterTemperature,
		DewpointTemperature: o.DewpointTemperature,
		Visibility:          o.Visibility,
		Tide:                o.Tide,
	}
	return json.Marshal(metObs)
}

// func (mo MeteorologicalObservation) String() string {
// 	return fmt.Sprintf("(MeteorologicalObservation datetime=%s, wind_direction=%v, wind_speed=%v, wind_gust=%v)",
// 		mo.Datetime, mo.WindDirection, mo.WindSpeed, mo.WindGust,
// 	)
// }

func Realtime(stationId string, dataset RealtimeDataset) ([]MeteorologicalObservation, error) {
	if valid := dataset.IsValid(); !valid {
		return nil, fmt.Errorf("unknown dataset '%s'", dataset)
	}
	url := fmt.Sprintf("%s/%s.%s", RealtimeURL, stationId, dataset)

	data, err := realtimeMeteorological(url)
	if err != nil {
		return nil, err
	}

	return data, nil
}
