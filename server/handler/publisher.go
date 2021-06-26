package handler

import (
	"net/http"

	"nipun.io/message_queue/appcontext"
	"nipun.io/message_queue/domain"
)

type PublisherHandler struct {
}

func NewPublisherHandler(dependencies *appcontext.Instance) *PublisherHandler {
	return &PublisherHandler{}
}

func (*PublisherHandler) PublishMessageAPI(w http.ResponseWriter, r *http.Request) {
	appcontext.Logger.Debug().Msg("Publisher API called")
	domain.WriteResponse(http.StatusOK, nil, w)
}
