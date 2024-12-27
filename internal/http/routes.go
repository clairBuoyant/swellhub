package http

import (
	"net/http"

	"github.com/clairBuoyant/swellhub/internal/app"
	"github.com/clairBuoyant/swellhub/web"
)

func NewRouter(app *app.Application) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/auth/me", AuthMe(app))

	mux.HandleFunc("/buoy/{stationID}", Buoy(app))
	mux.HandleFunc("/buoys", Buoys(app))

	mux.HandleFunc("/status", Status(app))

	mux.HandleFunc("/", web.SPAHandler())

	return mux
}
