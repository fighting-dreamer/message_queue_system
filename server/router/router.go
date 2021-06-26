package router

import (
	"github.com/gorilla/mux"
	"nipun.io/message_queue/appcontext"
)

func Router(dependencies *appcontext.Instance) *mux.Router {
	router := mux.NewRouter()

	handleSystemRoutes(dependencies, router)

	handlePublisherRoutes(dependencies, router)
	handleSubscriberRoutes(dependencies, router)
	handleQueueRoutes(dependencies, router)

	return router
}
