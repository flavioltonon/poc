package main

import (
	"log/slog"
	"os"

	"poc/shared/generic"
)

func main() {
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
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
	})

	logger := slog.New(handler)

	// time=2023-08-25T12:34:56.789Z level=DEBUG msg="logging at DEBUG level"
	logger.Debug("logging at DEBUG level")

	// time=2023-08-25T12:34:56.789Z level=INFO msg="logging at INFO level" some_group.foo=bar some_group.baz=1
	logger.Info("logging at INFO level", slog.Group("some_group", slog.String("foo", "bar"), slog.Int("baz", 1)))

	// time=2023-08-25T12:34:56.789Z level=WARN msg="logging at WARN level" foo=1h23m45s
	logger.Warn("logging at WARN level", slog.Duration("foo", generic.Duration))

	// time=2023-08-25T12:34:56.789Z level=ERROR msg="logging at ERROR level" error="some error"
	logger.Error("logging at ERROR level", slog.Any("error", generic.Error))
}
