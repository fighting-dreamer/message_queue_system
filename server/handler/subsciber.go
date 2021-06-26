package handler

import (
	"net/http"

	"nipun.io/message_queue/appcontext"
	"nipun.io/message_queue/domain"
)

type SubscriberHandler struct {
}

func NewsubscriberHandler(dependencies *appcontext.Instance) *SubscriberHandler {
	return &SubscriberHandler{}
}

func (sh *SubscriberHandler) RegisterSubscriberAPI(w http.ResponseWriter, r *http.Request) {
	appcontext.Logger.Debug().Msg("Subscriber Register API called")
	domain.WriteResponse(http.StatusOK, nil, w)
}

func (sh *SubscriberHandler) PollMessageAPI(w http.ResponseWriter, r *http.Request) {
	appcontext.Logger.Debug().Msg("Subscriber API called")
	domain.WriteResponse(http.StatusOK, nil, w)
}
