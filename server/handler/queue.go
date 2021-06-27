package handler

import (
	"io"
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
	// Parse the request
	bodyBytes, _ := io.ReadAll(r.Body)
	createQueueRequest := domain.CreateQueueRequest{}
	err := getBody(r.Context(), bodyBytes, &createQueueRequest)
	if err != nil {
		domain.WriteErrorResponse(http.StatusBadRequest, []string{JsonParseError.Error()}, w)
		return
	}

	// operate on request
	err = qh.QueueManager.CreateQueue(createQueueRequest)

	// send response
	if err != nil {
		domain.WriteErrorResponse(http.StatusInternalServerError, []string{err.Error()}, w)
		return
	}
	response := domain.CreateQueueResponse{
		Name: createQueueRequest.Name,
	}
	domain.WriteResponse(http.StatusOK, response, w)
}
