package main

import (
	"os"
	"runtime/debug"
	"sync"

	"github.com/clairBuoyant/swellhub/internal/app"
	"github.com/clairBuoyant/swellhub/internal/http"
	"github.com/clairBuoyant/swellhub/pkg/env"
	"github.com/clairBuoyant/swellhub/pkg/log"
)

func main() {
	cfg := &app.Config{
		Port: env.GetInt("PORT", 4000),
	}

	log := log.InitLoggerJSON()

	var wg sync.WaitGroup

	application := app.NewApplication(cfg, log, &wg)

	router := http.NewRouter(application)

	if err := application.ServeHTTP(router); err != nil {
		trace := string(debug.Stack())
		log.Error("could not start http server", "error", err, "trace", trace)
		os.Exit(1)
	}
}
