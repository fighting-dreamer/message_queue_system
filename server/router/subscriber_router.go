package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"nipun.io/message_queue/appcontext"
	"nipun.io/message_queue/server/handler"
)

func handleSubscriberRoutes(dependencies *appcontext.Instance, router *mux.Router) {
	subscriberHandler := handler.NewSubscriberHandler(dependencies)
	router.HandleFunc("/v1/subscribe/poll", subscriberHandler.PollMessageAPI).Methods(http.MethodPost)
	router.HandleFunc("/v1/subscribe/register", subscriberHandler.RegisterSubscriberAPI).Methods(http.MethodPost)
	router.HandleFunc("/v1/subscribe/message/ack", subscriberHandler.AckMessageAPI).Methods(http.MethodPost)
}
