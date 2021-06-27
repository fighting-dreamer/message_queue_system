package appcontext

import (
	"nipun.io/message_queue/config"
	"nipun.io/message_queue/logger"
)

func Init() {
	config.Load()
	LoadDependencies()

	logger.SetupLogger()
}
