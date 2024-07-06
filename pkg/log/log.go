package log

import (
	"log/slog"
	"os"
)

const (
	LevelTrace = slog.Level(-8)
	LevelDebug = slog.LevelDebug
	LevelInfo  = slog.LevelInfo
	LevelWarn  = slog.LevelWarn
	LevelError = slog.LevelError
	LevelFatal = slog.Level(12)
)

var LevelNames = map[slog.Leveler]string{
	LevelTrace: "TRACE",
	LevelFatal: "FATAL",
}

// Enums for Level
const (
	TraceLevel slog.Level = LevelTrace
	DebugLevel slog.Level = LevelDebug
	InfoLevel  slog.Level = LevelInfo
	WarnLevel  slog.Level = LevelWarn
	ErrorLevel slog.Level = LevelError
	FatalLevel slog.Level = LevelFatal
)

// InitLoggerJSON initializes global logger with full json format.
func InitLoggerJSON() *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: LevelTrace,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				levelLabel, exists := LevelNames[level]
				if !exists {
					levelLabel = level.String()
				}
				a.Value = slog.StringValue(levelLabel)
			}
			return a
		},
	}
	h := slog.NewJSONHandler(os.Stdout, opts)
	return slog.New(h)
}
