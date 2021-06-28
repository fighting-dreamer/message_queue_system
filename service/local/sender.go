package service

import (
	"nipun.io/message_queue/domain"
	"nipun.io/message_queue/logger"
	"nipun.io/message_queue/service"
)

const queueCap = 10

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

	// NOT IMPLEMENTING THE COMPLEX LOGIC FOR NOW
	//ackCounter := ss.SubscriberManager.GetAckCounter(queueName, request.SubscriberID)
	//unackCounter := ss.SubscriberManager.GetUnackCounter(queueName, request.SubscriberID)
	//fetchAsPerSubscriberRequest := request.MessageID + request.FetchCount
	//fetchFromMessageID := max(unackCounter, int64(request.MessageID))
	//fetchTillMessageID := min(ackCounter + queueCap, int64(fetchAsPerSubscriberRequest))

	for msgID := request.MessageID; msgID < (request.MessageID + request.FetchCount); msgID++ {
		message, err := ss.MessageBroker.GetMessage(queueRef, request.SubscriberID, int64(msgID))
		if err != nil {
			return []domain.Message{}, err
		}
		// err can be due to wrong mapping, if so, that can be ignored right now.
		ss.SubscriberManager.IncrementUnackCounter(queueName, request.SubscriberID)
		res = append(res, message)
	}
	logger.Logger.Debug().Msgf("unackCounter : %d, Ack Counter : %d", ss.SubscriberManager.GetUnackCounter(queueName, request.SubscriberID), ss.SubscriberManager.GetAckCounter(queueName, request.SubscriberID))
	logger.Logger.Debug().Msgf("%+v", res)

	return res, nil
}
