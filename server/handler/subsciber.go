package handler

import (
	"io"
	"net/http"

	"nipun.io/message_queue/appcontext"
	"nipun.io/message_queue/domain"
	"nipun.io/message_queue/logger"
	"nipun.io/message_queue/service"
)

type SubscriberHandler struct {
	SubscriberManager service.ISubscriberManager
	SenderService     service.ISenderService
}

func NewSubscriberHandler(dependencies *appcontext.Instance) *SubscriberHandler {
	return &SubscriberHandler{
		SubscriberManager: dependencies.SubscriberManager,
		SenderService:     dependencies.SenderService,
	}
}

func (sh *SubscriberHandler) RegisterSubscriberAPI(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Debug().Msg("Subscriber Register API called")
	// Parse the request
	bodyBytes, _ := io.ReadAll(r.Body)
	registerSubscriberRequest := domain.SubscriberRegisterRequest{}
	err := getBody(r.Context(), bodyBytes, &registerSubscriberRequest)
	if err != nil {
		domain.WriteErrorResponse(http.StatusBadRequest, []string{JsonParseError.Error()}, w)
		return
	}

	// operate on request
	err = sh.SubscriberManager.RegisterSubscriber(&registerSubscriberRequest)

	// response
	if err != nil {
		domain.WriteErrorResponse(http.StatusInternalServerError, []string{err.Error()}, w)
		return
	}
	response := domain.SubscriberRegisterResponse{
		SubscriberID: registerSubscriberRequest.SubscriberID,
		QueueName:    registerSubscriberRequest.QueueName,
	}
	domain.WriteResponse(http.StatusOK, response, w)
}

func (sh *SubscriberHandler) PollMessageAPI(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Debug().Msg("Subscriber API called")
	domain.WriteResponse(http.StatusOK, nil, w)
}
