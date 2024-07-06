package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/clairBuoyant/swellhub/internal/response"
)

const (
	defaultIdleTimeout    = time.Minute
	defaultReadTimeout    = 5 * time.Second
	defaultWriteTimeout   = 10 * time.Second
	defaultShutdownPeriod = 30 * time.Second
)

type Config struct {
	BaseURL string
	Port    int
}

type Application struct {
	Config    *Config
	Logger    *slog.Logger
	WaitGroup *sync.WaitGroup
}

func NewApplication(cfg *Config, logger *slog.Logger, wg *sync.WaitGroup) *Application {
	return &Application{Config: cfg, Logger: logger, WaitGroup: wg}
}

func (app *Application) ServeHTTP(handler http.Handler) error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.Config.Port),
		Handler:      handler,
		ErrorLog:     slog.NewLogLogger(app.Logger.Handler(), slog.LevelWarn),
		IdleTimeout:  defaultIdleTimeout,
		ReadTimeout:  defaultReadTimeout,
		WriteTimeout: defaultWriteTimeout,
	}

	shutdownErrorChan := make(chan error)

	go func() {
		quitChan := make(chan os.Signal, 1)
		signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)
		<-quitChan

		ctx, cancel := context.WithTimeout(context.Background(), defaultShutdownPeriod)
		defer cancel()

		shutdownErrorChan <- srv.Shutdown(ctx)
	}()

	app.Logger.Info("starting server", slog.Group("server", "addr", srv.Addr))

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownErrorChan
	if err != nil {
		return err
	}

	app.Logger.Info("stopped server", slog.Group("server", "addr", srv.Addr))

	app.WaitGroup.Wait()
	return nil
}

func (app *Application) ReportServerError(r *http.Request, err error) {
	var (
		message = err.Error()
		method  = r.Method
		url     = r.URL.String()
		trace   = string(debug.Stack())
	)

	requestAttrs := slog.Group("request", "method", method, "url", url)
	app.Logger.Error(message, requestAttrs, "trace", trace)
}

func (app *Application) ErrorMessage(w http.ResponseWriter, r *http.Request, status int, message string, headers http.Header) {
	message = strings.ToUpper(message[:1]) + message[1:]

	err := response.JSONWithHeaders(w, status, map[string]string{"Error": message}, headers)
	if err != nil {
		app.ReportServerError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *Application) ServerError(w http.ResponseWriter, r *http.Request, err error, code int) {
	message := err.Error()
	if code == 0 {
		code = http.StatusInternalServerError
		message = "server encountered a problem and could not process your request"
	}
	app.ReportServerError(r, err)

	app.ErrorMessage(w, r, code, message, nil)
}
