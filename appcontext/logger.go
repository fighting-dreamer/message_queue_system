package appcontext

import (
	"os"

	"github.com/rs/zerolog"
)

var Logger zerolog.Logger

func SetupLogger() {
	zerolog.SetGlobalLevel(getLogLevel())
	Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
}

func getLogLevel() zerolog.Level {

	level, err := zerolog.ParseLevel(LogLevel())

	if err != nil {
		return zerolog.InfoLevel
	}
	return level
}
