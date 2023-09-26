package main

import (
	"fmt"
	"os"
	"path/filepath"
	"poc/shared/generic"

	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.
		New(os.Stderr).
		Level(zerolog.DebugLevel).
		With().
		Caller().
		Time("time", generic.Time).
		Str("default", "banana").
		Logger()

	// Replace default marshal function for "caller" field
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		fullDir, file := filepath.Split(file)
		lastDir := filepath.Base(fullDir)
		return fmt.Sprintf("%s/%s:%d", lastDir, file, line)
	}

	// {"level":"debug","time":"2023-08-25T12:34:56Z","default":"banana","caller":"json/main.go:30","message":"logging at DEBUG level"}
	logger.Debug().Msg("logging at DEBUG level")

	// {"level":"info","time":"2023-08-25T12:34:56Z","default":"banana","map":{"foo":"bar","baz":1},"caller":"json/main.go:33","message":"logging at INFO level"}
	logger.Info().Dict("map", zerolog.Dict().Str("foo", "bar").Int("baz", 1)).Msg("logging at INFO level")

	// {"level":"warn","time":"2023-08-25T12:34:56Z","default":"banana","foo":"1h23m45s","caller":"json/main.go:36","message":"logging at WARN level"}
	logger.Warn().Str("foo", generic.Duration.String()).Msg("logging at WARN level")

	// {"level":"error","time":"2023-08-25T12:34:56Z","default":"banana","error":"some error","caller":"json/main.go:39","message":"logging at ERROR level"}
	logger.Error().Err(generic.Error).Msg("logging at ERROR level")
}
