package http

import (
	"net/http"

	"github.com/clairBuoyant/swellhub/internal/app"
)

func NewRouter(app *app.Application) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/status", Status(app))

	return mux
}
