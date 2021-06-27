package service

import (
	"nipun.io/message_queue/domain"
	"nipun.io/message_queue/logger"
	"nipun.io/message_queue/service"
)

type RecieverService struct {
	QueueManager  service.IQueueManager
	MessageBroker service.IMessageBrokerService
}

func (rs *RecieverService) EnqueueMessage(message *domain.Message) (*domain.Message, error) {
	logger.Logger.Debug().Msgf("RecieverService EnqueueMessage is called %+v", message.Value)
	queueName := message.Metadata.QueueName
	queueRef, err := rs.QueueManager.GetQueue(queueName)
	logger.Logger.Debug().Msgf("Got Queue : %s", queueRef.ID)
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
