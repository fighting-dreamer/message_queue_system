package service

import "nipun.io/message_queue/domain"

type IMessageStoreService interface {
	GetMessage()
	SetMessage(queueName string, message domain.Message) error
	DeleteMessage()
}
