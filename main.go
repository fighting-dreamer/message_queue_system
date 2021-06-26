package main

import (
	"nipun.io/message_queue/appcontext"
	"nipun.io/message_queue/server"
)

func main() {
	appcontext.Init()
	server.StartApiServer(appcontext.AppDependencies)
}
