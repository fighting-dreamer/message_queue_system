package handler

import (
	"net/http"

	"nipun.io/message_queue/appcontext"
	"nipun.io/message_queue/domain"
)

type QueueHandler struct {
}

func NewQueueHandler(dependencies *appcontext.Instance) *QueueHandler {
	return &QueueHandler{}
}

func (qh *QueueHandler) CreateQueueAPI(w http.ResponseWriter, r *http.Request) {
	appcontext.Logger.Debug().Msg("QueueHandler Create API called")
	domain.WriteResponse(http.StatusOK, nil, w)
}
