package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"nipun.io/message_queue/appcontext"
	"nipun.io/message_queue/server/handler"
)

func handleQueueRoutes(dependencies *appcontext.Instance, router *mux.Router) {
	queueHandler := handler.NewQueueHandler(dependencies)
	router.HandleFunc("/v1/queue/create", queueHandler.CreateQueueAPI).Methods(http.MethodPost)
}
