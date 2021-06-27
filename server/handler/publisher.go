package handler

import (
	"errors"
	"io"
	"net/http"

	"nipun.io/message_queue/appcontext"
	"nipun.io/message_queue/domain"
	"nipun.io/message_queue/logger"
	"nipun.io/message_queue/service"
)

var (
	JsonParseError = errors.New("JsonParseError")
)

type PublisherHandler struct {
	RecieverService service.IRecieverService
}

func NewPublisherHandler(dependencies *appcontext.Instance) *PublisherHandler {
	return &PublisherHandler{
		RecieverService: dependencies.RecieverService,
	}
}

func (ph *PublisherHandler) PublishMessageAPI(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Debug().Msg("Publisher API called")

	bodyBytes, _ := io.ReadAll(r.Body)
	message := domain.Message{}
	err := getBody(r.Context(), bodyBytes, &message)
	if err != nil {
		domain.WriteErrorResponse(http.StatusBadRequest, []string{JsonParseError.Error()}, w)
		return
	}

	ph.RecieverService.EnqueueMessage(message)
	domain.WriteResponse(http.StatusOK, nil, w)
}
