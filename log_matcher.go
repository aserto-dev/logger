package logger

import (
	api "github.com/aserto-dev/go-grpc/aserto/api/v1"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func LogLeverMatcher(level api.LogLevel) (zerolog.Level, error) {
	switch level {
	case api.LogLevel_LOG_LEVEL_TRACE:
		return zerolog.TraceLevel, nil
	case api.LogLevel_LOG_LEVEL_DEBUG:
		return zerolog.DebugLevel, nil
	case api.LogLevel_LOG_LEVEL_INFO:
		return zerolog.InfoLevel, nil
	case api.LogLevel_LOG_LEVEL_WARN:
		return zerolog.WarnLevel, nil
	case api.LogLevel_LOG_LEVEL_ERROR:
		return zerolog.ErrorLevel, nil
	case api.LogLevel_LOG_LEVEL_FATAL:
		return zerolog.FatalLevel, nil
	case api.LogLevel_LOG_LEVEL_PANIC:
		return zerolog.PanicLevel, nil
	default:
		return zerolog.NoLevel, status.Errorf(codes.InvalidArgument, "could not set log level to %s", level.String())
	}
}
