package service

import (
	"nipun.io/message_queue/domain"
	"nipun.io/message_queue/service"
)

type RecieverService struct {
	QueueManager  service.IQueueManager
	MessageBroker service.IMessageBrokerService
}

func (rs *RecieverService) EnqueueMessage(message *domain.Message) (*domain.Message, error) {
	queueName := message.Metadata.QueueName
	queueRef, err := rs.QueueManager.GetQueue(queueName)
	if err != nil {
		return nil, err
	}
	messageRef, err := rs.MessageBroker.SetMessage(queueRef, message)
	if err != nil {
		// send the original message
		return message, err
	}
	// send the modified message
	return messageRef, err
}
