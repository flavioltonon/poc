package main

import (
	"log/slog"
	"os"

	"poc/shared/generic"
)

func main() {
	handler := slog.
		NewJSONHandler(os.Stderr, &slog.HandlerOptions{
			// Adds a key with the source of the log
			AddSource: false,

			// Sets slog.LevelDebug as the minimum loggable level
			Level: slog.LevelDebug,

			// Replaces logged attributes
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				// Replaces default "time" key with a custom value
				if a.Key == slog.TimeKey {
					return slog.Time(slog.TimeKey, generic.Time)
				}

				// Replaces time.Duration keys with their string values
				if a.Value.Kind() == slog.KindDuration {
					return slog.String(a.Key, a.Value.Duration().String())
				}

				return a
			},
		}).
		// Adds default attributes to all logs
		WithAttrs([]slog.Attr{
			slog.String("default", "banana"),
		})

	logger := slog.New(handler)

	// {"time":"2023-08-25T12:34:56.789101112Z","level":"DEBUG","msg":"logging at DEBUG level","default":"banana"}
	logger.Debug("logging at DEBUG level")

	// {"time":"2023-08-25T12:34:56.789101112Z","level":"INFO","msg":"logging at INFO level","default":"banana","map":{"foo":"bar","baz":1}}
	logger.Info("logging at INFO level", slog.Group("map", slog.String("foo", "bar"), slog.Int("baz", 1)))

	// {"time":"2023-08-25T12:34:56.789101112Z","level":"WARN","msg":"logging at WARN level","default":"banana","foo":"1h23m45s"}
	logger.Warn("logging at WARN level", slog.Duration("foo", generic.Duration))

	// {"time":"2023-08-25T12:34:56.789101112Z","level":"ERROR","msg":"logging at ERROR level","default":"banana","error":"some error"}
	logger.Error("logging at ERROR level", slog.Any("error", generic.Error))
}
