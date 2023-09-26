package main

import (
	"os"
	"time"

	"poc/shared/generic"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		LevelKey:     "level",
		MessageKey:   "msg",
		CallerKey:    "caller",
		EncodeCaller: zapcore.ShortCallerEncoder,
		// Encode log level as its string value in uppercase
		EncodeLevel: zapcore.CapitalLevelEncoder,
		// Encode log timestamp as its RFC3339Nano string value
		EncodeTime: zapcore.RFC3339NanoTimeEncoder,
		// Encode time.Duration fields as their string values
		EncodeDuration: func(d time.Duration, pae zapcore.PrimitiveArrayEncoder) { pae.AppendString(d.String()) },
	})

	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stderr), zapcore.DebugLevel)

	logger := zap.New(core,
		zap.AddCaller(),
		zap.Fields(
			zap.String("default", "banana"),

			// Sets a "time" key with a custom value
			zap.Time("time", generic.Time),
		),
	)

	// {"level":"DEBUG","caller":"json/main.go:40","msg":"logging at DEBUG level","default":"banana","time":"2023-08-25T12:34:56.789101112Z"}
	logger.Debug("logging at DEBUG level")

	// {"level":"INFO","caller":"json/main.go:43","msg":"logging at INFO level","default":"banana","time":"2023-08-25T12:34:56.789101112Z","map":{"foo":"bar","baz":1}}
	logger.Info("logging at INFO level", zap.Dict("map", zap.String("foo", "bar"), zap.Int("baz", 1)))

	// {"level":"WARN","caller":"json/main.go:46","msg":"logging at WARN level","default":"banana","time":"2023-08-25T12:34:56.789101112Z","foo":"1h23m45s"}
	logger.Warn("logging at WARN level", zap.Duration("foo", generic.Duration))

	// {"level":"ERROR","caller":"json/main.go:49","msg":"logging at ERROR level","default":"banana","time":"2023-08-25T12:34:56.789101112Z","error":"some error"}
	logger.Error("logging at ERROR level", zap.Any("error", generic.Error))
}
