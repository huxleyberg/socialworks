package log

import (
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

// LegacyCompatibleLogger is a wrapper around zerolog.Logger that implements the legacy log.Logger interface in commons
type LegacyCompatibleLogger struct {
	zerolog.Logger
}

// Log is a method that satisfies the Log() for the log.Logger interface
func (lcl LegacyCompatibleLogger) Log(msg string) {
	if strings.Contains(strings.ToLower(msg), "error") {
		lcl.Error().Msg(msg)
		return
	}
	lcl.Logger.Log().Msg(msg)
}

// Logf is a method that satisfies the Logf() for the log.Logger interface
func (lcl LegacyCompatibleLogger) Logf(format string, v ...interface{}) {
	if strings.Contains(strings.ToLower(format), "error") {
		lcl.Error().Msgf(format, v...)
		return
	}
	lcl.Logger.Log().Msgf(format, v...)
}

// ZeroLogger is the global logger that will be used throughout the application
var ZeroLogger LegacyCompatibleLogger

// Init initializes the logger. It is automatically called when the package is imported
func init() {
	// Setting the global time format to UTC
	zerolog.TimestampFunc = func() time.Time {
		return time.Now().UTC()
	}

	// Configure zerolog for JSON output
	zerolog.TimeFieldFormat = time.RFC3339

	// Set log level from environment variable
	level, err := zerolog.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		level = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(level)
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	ZeroLogger.Logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Stack().Logger()
}
