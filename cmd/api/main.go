package main

import (
	"log/slog"
	"os"
	"runtime/debug"
	"sync"
)

type config struct {
	baseURL  string
	httpPort int
}

type application struct {
	config config
	logger *slog.Logger
	wg     sync.WaitGroup
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	err := run(logger)
	if err != nil {
		trace := string(debug.Stack())
		logger.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}

func run(logger *slog.Logger) error {
	var cfg config
	cfg.baseURL = "http://localhost" // env.GetInt("BASE_URL", "http://localhost")
	cfg.httpPort = 4000              // env.GetInt("HTTP_PORT", 4000)

	app := &application{
		config: cfg,
		logger: logger,
	}

	return app.serveHTTP()
}
