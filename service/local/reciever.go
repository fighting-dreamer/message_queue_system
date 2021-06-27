package service

import (
	"nipun.io/message_queue/domain"
	"nipun.io/message_queue/service"
)

type RecieverService struct {
	QueueManager  service.IQueueManager
	MessageBroker service.IMessageBrokerService
}

func (rs *RecieverService) EnqueueMessage(message domain.Message) {}
