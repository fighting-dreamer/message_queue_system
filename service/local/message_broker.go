package service

import (
	"nipun.io/message_queue/domain"
	"nipun.io/message_queue/logger"
	"nipun.io/message_queue/service"
)

type MessageBrokerService struct {
	MessageStoreService service.IMessageStoreService
	CallBackChan        chan *domain.Message
}

func (mbs *MessageBrokerService) SetMessage(queue *domain.Queue, message *domain.Message) (*domain.Message, error) {
	messageWithID, err := mbs.MessageStoreService.SetMessage(queue.ID, *message)
	if err != nil {
		return nil, err
	}
	mbs.CallSubscribers(queue, message)
	return &messageWithID, nil
}
func (mbs *MessageBrokerService) GetMessage(queue *domain.Queue, subscriberID string, messageID int) (domain.Message, error) {
	logger.Logger.Debug().Msg("Message Broker GetMessage")
	msg, err := mbs.MessageStoreService.GetMessage(int64(messageID))
	if err != nil {
		return domain.Message{}, err
	}
	return msg, nil
}

func (mbs *MessageBrokerService) CallSubscribers(queueRef *domain.Queue, message *domain.Message) error {
	logger.Logger.Debug().Msg("Message Broker callsubscribers")
	mbs.CallBackChan <- message
	return nil
}
