package service

import (
	"nipun.io/message_queue/service"
)

type SenderService struct {
	QueueManager      service.IQueueManager
	SubscriberManager service.ISubscriberManager
	MessageBroker     service.IMessageBrokerService
}

func (ss *SenderService) GetMessage() {}
