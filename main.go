package main

import (
	"nipun.io/message_queue/appcontext"
	"nipun.io/message_queue/logger"
	"nipun.io/message_queue/server"
	"nipun.io/message_queue/service"
)

func Start(worker service.ICallBackWorker) {
	logger.Logger.Debug().Msg("Started the Callback Worker")
	for messageRef := range worker.GetCallBackChan() {
		logger.Logger.Debug().Msgf("Got Message : %+v", *messageRef)
		go worker.CallSubscribers(messageRef)
	}
}

func main() {
	appcontext.Init()
	go Start(appcontext.AppDependencies.CallbackWorker)
	server.StartApiServer(appcontext.AppDependencies)
}
