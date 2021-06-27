package main

import (
	"nipun.io/message_queue/appcontext"
	"nipun.io/message_queue/server"
	local_service "nipun.io/message_queue/service/local"
)

func main() {
	appcontext.Init()
	go local_service.Start(appcontext.AppDependencies.CallbackWorker)
	server.StartApiServer(appcontext.AppDependencies)
}
