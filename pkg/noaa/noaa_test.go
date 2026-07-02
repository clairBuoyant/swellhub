package noaa

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRequestFailsFastWhenServerHangs(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		<-r.Context().Done() // never respond; unblocks when the client gives up
	}))
	defer srv.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	start := time.Now()
	_, _, err := request(ctx, srv.URL)
	if err == nil {
		t.Fatal("request against hanging server = nil error, want context deadline error")
	}
	if elapsed := time.Since(start); elapsed > 2*time.Second {
		t.Fatalf("request took %v to fail, want prompt cancellation", elapsed)
	}
}

func TestParseValuePreservesMissingAndZero(t *testing.T) {
	missing, err := parseValue("MM")
	if err != nil {
		t.Fatalf("parseValue(MM): %v", err)
	}
	if missing != nil {
		t.Fatalf("parseValue(MM) = %v, want nil", *missing)
	}

	zero, err := parseValue("0.0")
	if err != nil {
		t.Fatalf("parseValue(0.0): %v", err)
	}
	if zero == nil || *zero != 0 {
		t.Fatalf("parseValue(0.0) = %v, want pointer to zero", zero)
	}
}

func TestParseRecordPreservesMissingReadings(t *testing.T) {
	var observation MeteorologicalObservation
	record := []string{"2026 06 20 18 00 MM 0.0 MM 1.2 8 6 MM 1018.2 22.1 19.4 MM MM MM MM"}

	if err := parseRecordToStruct(record, &observation); err != nil {
		t.Fatalf("parseRecordToStruct: %v", err)
	}
	if observation.WindDirection != nil || observation.WindGust != nil || observation.WaveDirection != nil {
		t.Fatal("missing readings were reported as present")
	}
	if observation.WindSpeed == nil || *observation.WindSpeed != 0 {
		t.Fatalf("real zero wind speed = %v, want pointer to zero", observation.WindSpeed)
	}
	if observation.WaveHeight == nil || *observation.WaveHeight != 1.2 {
		t.Fatalf("wave height = %v, want 1.2", observation.WaveHeight)
	}
}
