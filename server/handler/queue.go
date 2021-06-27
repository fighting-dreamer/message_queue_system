package handler

import (
	"net/http"

	"nipun.io/message_queue/appcontext"
	"nipun.io/message_queue/domain"
	"nipun.io/message_queue/logger"
	"nipun.io/message_queue/service"
)

type QueueHandler struct {
	QueueManager service.IQueueManager
}

func NewQueueHandler(dependencies *appcontext.Instance) *QueueHandler {
	return &QueueHandler{
		QueueManager: dependencies.QueueManager,
	}
}

func (qh *QueueHandler) CreateQueueAPI(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Debug().Msg("QueueHandler Create API called")
	domain.WriteResponse(http.StatusOK, nil, w)
}
