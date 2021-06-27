package handler

import (
	"net/http"

	"nipun.io/message_queue/appcontext"
	"nipun.io/message_queue/domain"
	"nipun.io/message_queue/service"
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
	appcontext.Logger.Debug().Msg("Publisher API called")
	domain.WriteResponse(http.StatusOK, nil, w)
}
