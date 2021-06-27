package logger

import (
	"os"

	"github.com/rs/zerolog"
	"nipun.io/message_queue/config"
)

var Logger zerolog.Logger

func SetupLogger() {
	zerolog.SetGlobalLevel(getLogLevel())
	Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
}

func getLogLevel() zerolog.Level {

	level, err := zerolog.ParseLevel(config.LogLevel())

	if err != nil {
		return zerolog.InfoLevel
	}
	return level
}
