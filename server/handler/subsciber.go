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
	// Parse the request
	bodyBytes, _ := io.ReadAll(r.Body)
	subscriberPollRequest := domain.SubscriberPollRequest{}
	err := getBody(r.Context(), bodyBytes, &subscriberPollRequest)
	if err != nil {
		domain.WriteErrorResponse(http.StatusBadRequest, []string{JsonParseError.Error()}, w)
		return
	}

	messages, err := sh.SenderService.GetMessage(&subscriberPollRequest)
	if err != nil {
		domain.WriteErrorResponse(http.StatusInternalServerError, []string{err.Error()}, w)
		return
	}
	response := domain.SubscriberPollResponse{
		Messages: messages,
	}
	domain.WriteResponse(http.StatusOK, response, w)
}

func (sh *SubscriberHandler) AckMessageAPI(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Debug().Msg("Subscriber AckMessageAPI API called")
	// Parse the request
	bodyBytes, _ := io.ReadAll(r.Body)
	ackMessageRequest := domain.AckMessageRequest{}
	err := getBody(r.Context(), bodyBytes, &ackMessageRequest)
	if err != nil {
		domain.WriteErrorResponse(http.StatusBadRequest, []string{JsonParseError.Error()}, w)
		return
	}
	// silent API
	err = sh.SubscriberManager.IncrementAckCounter(ackMessageRequest.QueueName, ackMessageRequest.SubscriberID)
	if err != nil {
		domain.WriteErrorResponse(http.StatusBadRequest, []string{err.Error()}, w)
		return
	}
	response := ackMessageRequest
	domain.WriteResponse(http.StatusOK, response, w)
}
