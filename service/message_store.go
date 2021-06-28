package service

import "nipun.io/message_queue/domain"

type IMessageStoreService interface {
	GetMessage(messageID int64) (domain.Message, error)
	SetMessage(queueName string, message domain.Message) (domain.Message, error)
}
