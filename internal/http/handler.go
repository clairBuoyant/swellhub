package http

import (
	"net/http"

	"github.com/clairBuoyant/swellhub/internal/errors"
	"github.com/clairBuoyant/swellhub/internal/response"
)

type Application interface {
	ServerError(w http.ResponseWriter, r *http.Request, err error)
}

func Status(app Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"status": "OK",
		}
		if err := response.JSON(w, http.StatusOK, data); err != nil {
			app.ServerError(w, r, errors.NewAppError(http.StatusInternalServerError, "failed to encode response", err))
		}
	}
}
