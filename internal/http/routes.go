package http

import (
	"net/http"

	"github.com/clairBuoyant/swellhub/gen/clairbuoyant/spot/v1/spotv1connect"
	"github.com/clairBuoyant/swellhub/internal/app"
	"github.com/clairBuoyant/swellhub/internal/spotsvc"
	"github.com/clairBuoyant/swellhub/web"
)

func NewRouter(app *app.Application) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/auth/me", AuthMe(app))

	mux.HandleFunc("/buoy/{stationID}", Buoy(app))
	mux.HandleFunc("/buoys", Buoys(app))

	// Connect SpotService (clairbuoyant.spot.v1). Coexists with the legacy REST
	// buoy endpoints on the same mux; matched before the SPA catch-all.
	spotPath, spotHandler := spotv1connect.NewSpotServiceHandler(spotsvc.New(app.Logger))
	mux.Handle(spotPath, spotHandler)

	mux.HandleFunc("/status", Status(app))

	mux.HandleFunc("/", web.SPAHandler())

	return mux
}
