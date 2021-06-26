package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"nipun.io/message_queue/appcontext"
	"nipun.io/message_queue/server/handler"
)

func handlePublisherRoutes(dependencies *appcontext.Instance, router *mux.Router) {
	publisherHandler := handler.NewPublisherHandler(dependencies)
	router.HandleFunc("/v1/publish", publisherHandler.PublishMessageAPI).Methods(http.MethodPost)
}
