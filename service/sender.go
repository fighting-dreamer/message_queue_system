package service

import "nipun.io/message_queue/domain"

type ISenderService interface {
	GetMessage(request *domain.SubscriberPollRequest) ([]domain.Message, error)
}
