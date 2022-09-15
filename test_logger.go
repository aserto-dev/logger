package logger

import (
	"io"

	"github.com/rs/zerolog"
)

// TestLogger creates a logger useful for unit tests
// it has a trace level
func TestLogger(logOutput io.Writer) *zerolog.Logger {
	cfg := Config{}
	cfg.LogLevel = "trace"
	cfg.LogLevelParsed = zerolog.TraceLevel

	logger, err := NewLogger(logOutput, logOutput, &cfg)
	if err != nil {
		panic(err)
	}

	return logger
}
