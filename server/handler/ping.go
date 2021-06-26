package handler

import (
	"encoding/json"
	"net/http"

	"nipun.io/message_queue/appcontext"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	appcontext.Logger.Debug().Msg("Pinging API Server")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"success": "pong",
	})
}
