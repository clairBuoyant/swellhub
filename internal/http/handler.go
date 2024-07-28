package http

import (
	"fmt"
	"net/http"

	"github.com/clairBuoyant/swellhub/internal/errors"
	"github.com/clairBuoyant/swellhub/internal/response"
	"github.com/clairBuoyant/swellhub/pkg/noaa"
)

type application interface {
	ServerError(w http.ResponseWriter, r *http.Request, err error, code int)
}

func AuthMe(app application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type Me struct {
			FirstName string `json:"firstName"`
		}

		// TODO(@kylejb): provide user information
		if err := response.JSON(w, http.StatusOK, Me{
			FirstName: "UserFirstNameGoesHere (TESTING)",
		}); err != nil {
			app.ServerError(w, r, errors.NewAppError(http.StatusInternalServerError, "failed to encode response", err), 0)
		}
	}
}

func Buoy(app application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dataset := noaa.RealtimeDataset(r.URL.Query().Get("dataset"))
		if dataset == "" {
			dataset = noaa.TXT
		}

		if valid := dataset.IsValid(); !valid {
			app.ServerError(w, r, errors.NewAppError(http.StatusBadRequest, fmt.Sprintf("unknown dataset %s", dataset), nil), http.StatusBadRequest)
			return
		}

		// TODO(@kylejb): add validation for stationID too
		data, err := noaa.Realtime(r.PathValue("stationID"), dataset)
		if err != nil {
			// TODO(@kylejb): dynamically provide http status code
			app.ServerError(w, r, errors.NewAppError(http.StatusNotAcceptable, "failed to encode response", err), http.StatusNotAcceptable)
			return
		}
		if err := response.JSON(w, http.StatusOK, data); err != nil {
			app.ServerError(w, r, errors.NewAppError(http.StatusInternalServerError, "failed to encode response", err), 0)
		}
	}
}

func Buoys(app application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := noaa.ActiveStations()
		if err != nil {
			app.ServerError(w, r, errors.NewAppError(http.StatusBadRequest, "failed to retrieve active stations from NDBC", err), http.StatusBadRequest)
			return
		}

		if err := response.JSON(w, http.StatusOK, data); err != nil {
			app.ServerError(w, r, errors.NewAppError(http.StatusInternalServerError, "failed to encode response", err), 0)
		}
	}
}

func Status(app application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"status": "OK",
		}
		if err := response.JSON(w, http.StatusOK, data); err != nil {
			app.ServerError(w, r, errors.NewAppError(http.StatusInternalServerError, "failed to encode response", err), 0)
		}
	}
}
