package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"nipun.io/message_queue/appcontext"
	"nipun.io/message_queue/server/handler"
)

func handleSystemRoutes(dependencies *appcontext.Instance, router *mux.Router) {
	router.HandleFunc("/ping", handler.PingHandler).
		Methods(http.MethodGet)
}
