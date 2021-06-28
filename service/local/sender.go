package service

import (
	"nipun.io/message_queue/domain"
	"nipun.io/message_queue/logger"
	"nipun.io/message_queue/service"
)

type SenderService struct {
	QueueManager      service.IQueueManager
	SubscriberManager service.ISubscriberManager
	MessageBroker     service.IMessageBrokerService
}

func (ss *SenderService) GetMessage(request *domain.SubscriberPollRequest) ([]domain.Message, error) {
	queueName := ss.SubscriberManager.GetSubscriberQueueName(request.SubscriberID)
	logger.Logger.Debug().Msgf("SenderService GetMEssage subscriber : %s, Queue : %s", request.SubscriberID, queueName)
	// we can check if this is a valid queue
	queueRef, err := ss.QueueManager.GetQueue(queueName)
	if err != nil {
		logger.Logger.Debug().Msg("SenderService GetQueue from Queuemanager failed")
		return []domain.Message{}, err
	}

	res := []domain.Message{}

	for i := 0; i < request.FetchCount; i++ {
		msgID := request.MessageID + i
		message, err := ss.MessageBroker.GetMessage(queueRef, request.SubscriberID, msgID)
		if err != nil {
			// either all batch messages go through, OR none
			return []domain.Message{}, err
		}
		res = append(res, message)
	}

	logger.Logger.Debug().Msgf("%+v", res)

	return res, nil
}
