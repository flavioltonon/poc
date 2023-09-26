package main

import (
	"os"
	"poc/shared/generic"

	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.
		New(os.Stderr).
		Level(zerolog.DebugLevel).
		With().
		Time("time", generic.Time).
		Str("default", "banana").
		Logger()

	// {"level":"debug","time":"2023-08-25T12:34:56Z","default":"banana","message":"logging at DEBUG level"}
	logger.Debug().Msg("logging at DEBUG level")

	// {"level":"info","time":"2023-08-25T12:34:56Z","default":"banana","map":{"foo":"bar","baz":1},"message":"logging at INFO level"}
	logger.Info().Dict("map", zerolog.Dict().Str("foo", "bar").Int("baz", 1)).Msg("logging at INFO level")

	// {"level":"warn","time":"2023-08-25T12:34:56Z","default":"banana","foo":5025000,"message":"logging at WARN level"}
	logger.Warn().Dur("foo", generic.Duration).Msg("logging at WARN level")

	// {"level":"error","time":"2023-08-25T12:34:56Z","default":"banana","error":"some error","message":"logging at ERROR level"}
	logger.Error().Err(generic.Error).Msg("logging at ERROR level")
}
