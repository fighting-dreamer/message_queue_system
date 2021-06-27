package service

import "nipun.io/message_queue/domain"

type IMessageBrokerService interface {
	SetMessage(queue *domain.Queue, message *domain.Message) (*domain.Message, error)
	GetMessage(queue *domain.Queue, subscriberID string, messageID int) domain.Message
	DeleteMessage()
	CallSubscribers(queue *domain.Queue, message *domain.Message) error
}
