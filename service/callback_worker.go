package service

import "nipun.io/message_queue/domain"

type ICallBackWorker interface {
	GetCallBackChan() chan *domain.Message
	CallSubscribers(message *domain.Message) error
}
