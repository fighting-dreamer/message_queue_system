package service

import "nipun.io/message_queue/domain"

type IMessageBrokerService interface {
	SetMessage(queue *domain.Queue, message *domain.Message) (*domain.Message, error)
	GetMessage()
	DeleteMessage()
	CallSubscribers(queue *domain.Queue, message *domain.Message) error
}
