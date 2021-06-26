package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"nipun.io/message_queue/appcontext"
	"nipun.io/message_queue/server/handler"
)

func handleSubscriberRoutes(dependencies *appcontext.Instance, router *mux.Router) {
	subscriberHandler := handler.NewsubscriberHandler(dependencies)
	router.HandleFunc("/v1/subscribe", subscriberHandler.PollMessageAPI).Methods(http.MethodPost)

}
